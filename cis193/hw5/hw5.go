// Homework 5: Goroutines
// Due March 3, 2017 at 11:59pm
package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

func main() {
	// Feel free to use the main function for testing your functions
	hello := make(chan string, 5)
	hello <- "Hello world"
	hello <- "Привет мир"
	hello <- "Привіт Світ"
	hello <- "Witaj świecie"
	close(hello)
	fmt.Println(<-hello)
	for greeting := range hello {
		fmt.Println(greeting)
	}
}

// Filter copies values from the input channel into an output channel that match the filter function p
// The function p determines whether an int from the input channel c is sent on the output channel
func Filter(c <-chan int, p func(int) bool) <-chan int {
	buffCh := make(chan int, 9)
	go func() {
		for i := range c {
			if p(i) {
				buffCh <- i
			}
		}
		close(buffCh)
	}()

	return buffCh
}

// Result is a type representing a single result with its index from a slice
type Result struct {
	index  int
	result string
}

// ConcurrentRetry runs all the tasks concurrently and sends the output in a Result channel
//
// concurrent is the limit on the number of tasks running in parallel. Your
// solution must not run more than `concurrent` number of tasks in parallel.
//
// retry is the number of times that the task should be attempted. If a task
// returns an error, the function should be retried immediately up to `retry`
// times. Only send the results of a task into the output channel if it does not error.
//
// Multiple instances of ConcurrentRetry should be able to run simultaneously
// without interfering with one another, so global variables should not be used.
// The function must return the channel without waiting for the tasks to
// execute, and all results should be sent on the output channel. Once all tasks
// have been completed, close the channel.
func ConcurrentRetry(tasks []func() (string, error), concurrent int, retry int) <-chan Result {
	type taskInfo struct {
		idx  int
		task func() (string, error)
	}
	// result channel (return)
	res := make(chan Result)
	// channel for tasks
	taskList := make(chan taskInfo, len(tasks))
	// task waiting group
	var taskWG sync.WaitGroup

	gopherWork := func(i int, recTaskCh <-chan taskInfo) {
		for t := range recTaskCh {
			defer taskWG.Done()
			log.Printf("START\ttask#%d worker: gopher#%d\n", t.idx, i)

			for j := 0; j < retry; j++ {
				s, err := t.task()
				if err == nil {
					res <- Result{t.idx, s}
					log.Printf("END  \ttask#%d worker: gopher#%d\n", t.idx, i)
					break
				}
			}
		}
		log.Printf("GOHOME\tgopher#%v\n", i)
	}

	// we will have concurrent number of gophers work for us
	go func() {
		for i := 0; i < concurrent; i++ {
			go gopherWork(i, taskList)
		}
	}()

	// push tasks to channel
	for i, t := range tasks {
		log.Printf("Add  \ttask#%d\n", i)
		taskWG.Add(1)
		taskList <- taskInfo{i, t}
	}
	close(taskList)

	// wait ... close channel ...
	go func() {
		taskWG.Wait() // blocks here until WaitGroup counter is 0
		close(res)
	}()

	return res
}

// Task is an interface for types that process integers
type Task interface {
	Execute(int) (int, error)
}

// Fastest returns the result of the fastest running task
// Fastest accepts any number of Task structs. If no tasks are submitted to
// Fastest(), it should return an error.
// You should return the result of a Task even if it errors.
// Do not leave any pending goroutines. Make sure all goroutines are cleaned up
// properly and any synchronizing mechanisms closed.
func Fastest(input int, tasks ...Task) (int, error) {
	type taskResult struct {
		o int
		e error
	}

	if len(tasks) == 0 {
		return 0, errors.New("No input task.")
	}

	abort := make(chan struct{})
	done := make(chan struct{})
	res := make(chan taskResult)

	// for interrupt operations
	switcher := func() bool {
		select {
		case <-abort:
			return true
		default:
			return false
		}
	}

	for _, t := range tasks {
		go func(tk Task) {
			if switcher() {
				return
			}
			out, err := tk.Execute(input)
			if !switcher() {
				res <- taskResult{out, err}
				close(abort)
			}
		}(t)
	}

	go func() {
		for !switcher() {
		}
		close(res) // make sure that channel are closed
	}()

	var first taskResult
	go func() {
		first = <-res
		for r := range res {
			fmt.Println(r)
		}
		close(done)
	}()

	<-done
	return first.o, first.e
}

// MapReduce takes any number of tasks, and feeds their results through reduce
// If no tasks are supplied, return an error.
// If any of the tasks error during their execution, return an error immediately.
// Once all tasks have completed successfully, return the value of reduce on
// their results in any order.
// Do not leave any pending goroutines. Make sure all goroutines are cleaned up
// properly and any synchronizing mechanisms closed.
func MapReduce(input int, reduce func(results []int) int, tasks ...Task) (int, error) {
	// 1. tasks => taskCh (chan Task)
	// 2. Limited number of gophers work on tasks
	//   - output => output channel
	// 3. goroutine pull data from output channel
	//   - get a slice of output
	// 4. feed reduce func with the slice
	// * if any error occured, abort & clean goroutines
	//   - taskCh = drain
	//   - outCh = close & drain
	type taskResult struct {
		o int
		e error
	}
	// for abort
	abort := make(chan struct{})

	abortSwicher := func() bool {
		select {
		case <-abort:
			return true
		default:
			return false
		}
	}

	// step 1
	taskCh := make(chan Task)
	go func() {
		for _, t := range tasks {
			taskCh <- t
		}
		close(taskCh)
	}()
	// step 2
	// in order to close channel(outCh), use WaitGroup
	var wg sync.WaitGroup
	workerNum := 4
	outCh := make(chan taskResult)
	errCh := make(chan taskResult, 10) // blocked without a buffer
	for i := 0; i < workerNum; i++ {
		go func() {
			for t := range taskCh {
				if abortSwicher() {
					return
				}
				wg.Add(1)
				o, e := t.Execute(input)
				if e != nil && !abortSwicher() {
					errCh <- taskResult{o, e}
					close(abort)
				} else {
					outCh <- taskResult{o, e}
				}
				wg.Done()
			}
		}()
	}

	go func() {
		wg.Wait()
		close(outCh)
		close(errCh)
	}()

	// step 3
	res := []int{}
	done := make(chan struct{})
	go func() {
		for o := range outCh {
			if abortSwicher() {
				for range outCh {
				} // drain outCh
				for range taskCh {
				} // drain taskCh
				break
			}
			res = append(res, o.o)
		}
		close(done)
	}()
	// step 4
	<-done

	if abortSwicher() {
		o := <-errCh
		for range errCh {
		} // drain errCh
		return o.o, o.e
	} else {
		return reduce(res), nil
	}
}
