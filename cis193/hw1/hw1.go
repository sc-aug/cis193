// Homework 1: Finger Exercises
// Due January 31, 2017 at 11:59pm
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// Feel free to use the main function for testing your functions
	fmt.Println("Hello, دنيا!")
}

// ParsePhone parses a string of numbers into the format (123) 456-7890.
// This function should handle any number of extraneous spaces and dashes.
// All inputs will have 10 numbers and maybe extra spaces and dashes.
// For example, ParsePhone("123-456-7890") => "(123) 456-7890"
//              ParsePhone("1 2 3 4 5 6 7 8 9 0") => "(123) 456-7890"
func ParsePhone(phone string) string {
	var num = make([]rune, 10)
	var cnt = 0
	for _, n := range phone {
		if cnt >= 10 {
			break
		}
		if unicode.IsDigit(n) {
			num[cnt], cnt = n, cnt+1
		}
	}
	l := string(num[0:3])
	m := string(num[3:6])
	r := string(num[6:10])

	return fmt.Sprintf("(%s) %s-%s", l, m, r)
}

// Anagram tests whether the two strings are anagrams of each other.
// This function is NOT case sensitive and should handle UTF-8
func Anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)
	map1 := make(map[rune]int)
	map2 := make(map[rune]int)

	for _, c := range s1 {
		map1[c]++
	}

	for _, c := range s2 {
		map2[c]++
	}

	for key, cnt := range map1 {
		if map2[key] != cnt {
			return false
		}
	}

	return true
}

// FindEvens filters out all odd numbers from input slice.
// Result should retain the same ordering as the input.
func FindEvens(e []int) []int {
	res := make([]int, 0)
	for _, n := range e {
		if n&1 == 0 {
			res = append(res, n)
		}
	}
	return res
}

// SliceProduct returns the product of all elements in the slice.
// For example, SliceProduct([]int{1, 2, 3}) => 6
func SliceProduct(e []int) int {
	res := 1
	for _, n := range e {
		res *= n
	}
	return res
}

// Unique finds all distinct elements in the input array.
// Result should retain the same ordering as the input.
func Unique(e []int) []int {
	res := make([]int, 0)
	m := make(map[int]bool)
	for _, n := range e {
		_, contains := m[n]
		if contains {
			m[n] = false
		} else {
			m[n] = true
		}
	}

	for key, flag := range m {
		if flag {
			res = append(res, key)
		}
	}
	return res
}

// InvertMap inverts a mapping of strings to ints into a mapping of ints to strings.
// Each value should become a key, and the original key will become the corresponding value.
// For this function, you can assume each value is unique.
func InvertMap(kv map[string]int) map[int]string {
	res := make(map[int]string)
	for key, val := range kv {
		res[val] = key
	}
	return res
}

// TopCharacters finds characters that appear more than k times in the string.
// The result is the set of characters along with their occurrences.
// This function MUST handle UTF-8 characters.
func TopCharacters(s string, k int) map[rune]int {
	res := make(map[rune]int)
	for _, c := range s {
		res[c]++
	}
	for key, n := range res {
		if n <= k {
			delete(res, key)
		}
	}
	return res
}
