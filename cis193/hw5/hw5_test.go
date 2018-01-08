package main

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"
)

// Test Filter
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

// Test ConcurrentRetry
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

// Test Fastest
type test struct {
	val int
}

func (t test) Execute(n int) (int, error) {
	time.Sleep(time.Duration(100+n*t.val) * time.Millisecond)
	if t.val&1 == 0 {
		return 0, errors.New(fmt.Sprintf("error intended for: %v", t.val))
	}
	return n * t.val, nil
}

func TestFastest(t *testing.T) {
	sample := []test{
		test{0}, // with error
		test{1},
		test{2}, // with error
		test{3},
	}

	o, e := Fastest(3, sample[0], sample[1], sample[2], sample[3])
	fmt.Printf("Fastest Result: %v %v\n", o, e)

}

// Test MapReduce
func TestMapReduceNoError(t *testing.T) {
	sample := []test{
		test{1},
		test{3},
		test{5},
		test{7},
	}

	sum := func(nums []int) int {
		s := 0
		for _, n := range nums {
			s += n
		}
		return s
	}

	o, e := MapReduce(3, sum, sample[0], sample[1], sample[2], sample[3])
	fmt.Printf("MapReduce Result: %v %v\n", o, e)
}

// Test MapReduce
func TestMapReduceWithError(t *testing.T) {
	sample := []test{
		test{1},
		test{3},
		test{5},
		test{8},
	}

	sum := func(nums []int) int {
		s := 0
		for _, n := range nums {
			s += n
		}
		return s
	}

	o, e := MapReduce(3, sum, sample[0], sample[1], sample[2], sample[3])
	fmt.Printf("MapReduce Result: %v %v\n", o, e)
}
