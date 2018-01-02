package main

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestFilter(t *testing.T) {
	ch := make(chan int)
	sample := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expect := []int{2, 4, 6, 8}
	res := []int{}
	isEven := func(i int) bool { return i&1 == 0 }
	// goroutine put sample input channel
	go func() {
		for _, n := range sample {
			ch <- n
		}
		close(ch)
	}()
	// get channel from Filter function to receive data
	for i := range Filter(ch, isEven) {
		res = append(res, i)
	}
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Error: Filter. Expecting %v, get %v", expect, res)
	}
}

func TestConcurrentRetry(t *testing.T) {
	tasks := []func() (string, error){
		func() (string, error) {
			time.Sleep(2 * time.Second)
			return "sample1", nil
		},
		func() (string, error) {
			time.Sleep(3 * time.Second)
			return "sample2", nil
		},
		func() (string, error) {
			time.Sleep(time.Second)
			return "sample3", nil
		},
		func() (string, error) {
			time.Sleep(time.Second)
			return "sample4", nil
		},
	}

	resCh := ConcurrentRetry(tasks, 2, 2) // concurrent = 2, retry = 2
	// result
	for r := range resCh {
		fmt.Println("get result", r)
	}
}
