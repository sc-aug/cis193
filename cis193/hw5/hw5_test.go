package main

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	ch := make(chan int)
	sample := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expect := []int{2, 4, 6, 8}
	res := []int{}
	go func() {
		for _, n := range sample {
			ch <- n
		}
		close(ch)
	}()
	isEven := func(i int) bool {
		return i&1 == 0
	}
	for i := range Filter(ch, isEven) {
		res = append(res, i)
	}
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Error: Filter. Expecting %v, get %v", expect, res)
	}
}
