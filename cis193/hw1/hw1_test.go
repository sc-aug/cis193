package main

import (
	"reflect"
	"testing"
)

func TestParsePhone(t *testing.T) {
	sample := []string{"123-456-7890", "1 2 3 4 5 6 7 8 9 0", "(123)4567890"}
	res := []string{"(123) 456-7890", "(123) 456-7890", "(123) 456-7890"}
	for i, s := range sample {
		if ParsePhone(s) != res[i] {
			t.Errorf(`ParsePhone(%v) != %v`, s, res[i])
		}
	}
}

func TestAnagram(t *testing.T) {
	sample1 := []string{"cat", "1k98", "世界hello"}
	sample2 := []string{"act", "k189", "Hello世界"}
	res := []bool{true, true, true}
	for i, r := range res {
		if r != Anagram(sample1[i], sample2[i]) {
			t.Errorf(`Anagram(%v, %v) != %v`, sample1[i], sample2[i], r)
		}
	}
}

func TestFindEvens(t *testing.T) {
	sample := []int{1, 2, 3, 4, 5, 12, 15, 16, 17, 19, 23, 35, 63, 123, 135, 1234}
	res := []int{2, 4, 12, 16, 1234}
	equal := reflect.DeepEqual(FindEvens(sample), res)
	if !equal {
		t.Errorf(`FindEvens(%v) = %v != %v`, sample, FindEvens(sample), res)
	}
}

func TestSliceProduct(t *testing.T) {
	sample := []int{1, 3, 4}
	res := 12
	if res != SliceProduct(sample) {
		t.Errorf(`SliceProduct(%v) != %v`, sample, res)
	}
}

func TestUnique(t *testing.T) {
	sample := []int{1, 2, 11, 1, 5, 2, 5, 2, 1, 8, 5, 2, 6, 4, 2, 1, 6, 3, 7, 7, 4, 1, 7, 9}
	res := map[int]bool{
		11: true,
		9:  true,
		8:  true,
		3:  true,
	}
	output := Unique(sample)
	// same size
	if len(output) != len(res) {
		t.Errorf(`Function Unique() Error: size not match`, Unique(sample), res)
	}
	// same numbers
	for _, num := range output {
		if !res[num] {
			t.Errorf(`Function Unique() Error: numbers not match`, Unique(sample), res[num])
			break
		}
	}
}

func TestInvertMap(t *testing.T) {
	sample := map[string]int{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}
	res := map[int]string{
		2138: "r",
		1908: "gri",
		912:  "adg",
		3711: "rsc",
	}

	equal := reflect.DeepEqual(InvertMap(sample), res)
	if !equal {
		t.Errorf(`InvertMap(%v) != %v\n`, sample, res)
	}
}

func TestTopCharacters(t *testing.T) {
	sample_param1 := "abcdaabbbccccddddd的的的的世世世世界界界"
	sample_param2 := 3
	res := map[rune]int{
		'c': 5,
		'd': 6,
		'的': 4,
		'世': 4,
		'b': 4,
	}

	equal := reflect.DeepEqual(TopCharacters(sample_param1, sample_param2), res)
	if !equal {
		t.Errorf(`TopCharacters(%v, %v) != %v\n`, sample_param1, sample_param2, res)
	}
}
