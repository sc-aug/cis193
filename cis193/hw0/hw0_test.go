package main

import "testing"

func TestFizzbuzz(t *testing.T) {
	sample := []int{30, 5, 6, 16, 17, 18, 34, 33}
	res := []string{"FizzBuzz", "Buzz", "Fizz", "", "", "Fizz", "", "Fizz"}
	for i, s := range sample {
		if Fizzbuzz(s) != res[i] {
			t.Errorf(`Fizzbuzz(%v) != "%v"\n`, s, res[i])
		}
	}
}

func TestIsPrime(t *testing.T) {
	sample := []int{1, 2, 3, 4, 5, 6, 7, 111, 112, 113, 97, 91, 87, 88}
	res := []bool{false, true, true, false, true, false, true, false, false, true, true, false, false, false}
	for i, s := range sample {
		if IsPrime(s) != res[i] {
			t.Errorf(`IsPrime(%v) != %v\n`, s, res[i])
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	sample := []string{"", "f", "12321", "222", "22", "df919fd", "df99fd", "22432"}
	res := []bool{true, true, true, true, true, true, true, false}
	for i, s := range sample {
		if IsPalindrome(s) != res[i] {
			t.Errorf(`IsPalindrome(%v) != %v\n`, s, res[i])
		}
	}
}
