package main

import (
	"sort"
	"testing"
)

func TestNewPerson(t *testing.T) {
	sample := []Person{
		Person{1, "Fei", "Zhang"},
		Person{2, "Yu", "Guan"},
		Person{3, "Bei", "Liu"},
	}
	for _, s := range sample {
		if p := NewPerson(s.FirstName, s.LastName); *p != s {
			t.Errorf(`NewPerson Error: %v != %v`, p, s)
		}
	}
}

func TestSortInterface(t *testing.T) {
	sample := PersonSlice{
		&Person{1, "Fei", "Zhang"},
		&Person{2, "Yu", "Guan"},
		&Person{3, "Bei", "Liu"},
		&Person{4, "Bei", "Liu"},
		&Person{5, "Bei", "Bei"},
	}
	sort.Sort(sample)
	expect := []Person{
		Person{5, "Bei", "Bei"},
		Person{3, "Bei", "Liu"},
		Person{4, "Bei", "Liu"},
		Person{1, "Fei", "Zhang"},
		Person{2, "Yu", "Guan"},
	}
	for i, p := range sample {
		if *p != expect[i] {
			t.Errorf(`SortInterface Error: index=%v : %v != %v`, i, p, expect[i])
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	p1 := &Person{1, "a", "b"}
	p2 := &Person{2, "c", "d"}

	var sample = []struct {
		input  PersonSlice
		expect bool
	}{
		{PersonSlice{}, true},
		{PersonSlice{p1}, true},
		{PersonSlice{p1, p2}, false},
		{PersonSlice{p1, p2, p1}, true},
	}
	for _, s := range sample {
		if res := IsPalindrome(s.input); res != s.expect {
			t.Errorf(`IsPalindrome(%v) get %v, expecting %v`, s.input, s.expect)
		}
	}
}

func TestFold(t *testing.T) {
	add := func(x, y int) int { return x + y }
	prod := func(x, y int) int { return x * y }
	s := []int{1, 2, 3, 4}
	if Fold(s, 0, add) != 10 {
		t.Errorf(`Fold Error. f = %v, s = %v`, add, s)
	}
	if Fold(s, 1, prod) != 24 {
		t.Errorf(`Fold Error. f = %v, s = %v`, prod, s)
	}
}
