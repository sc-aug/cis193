// Homework 0: Hello Go!
// Due January 24, 2017 at 11:59pm
package main

import (
	"fmt"
	"strings"
)

func main() {
	// Feel free to use the main function for testing your functions
	fmt.Println("Hello, दुनिया!")
}

// Fizzbuzz is a classic introductory programming problem.
// If n is divisible by 3, return "Fizz"
// If n is divisible by 5, return "Buzz"
// If n is divisible by 3 and 5, return "FizzBuzz"
// Otherwise, return the empty string
func Fizzbuzz(n int) string {
	if n%15 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	} else {
		return ""
	}
}

// IsPrime checks if the number is prime.
// You may use any prime algorithm, but you may NOT use the standard library.
func IsPrime(n int) bool {
	if n <= 3 {
		return n >= 2
	}
	if n&1 == 0 {
		return false
	}
	for div := 3; div*div <= n; div += 2 {
		if n%div == 0 {
			return false
		}
	}
	return true
}

// IsPalindrome checks if the string is a palindrome.
// A palindrome is a string that reads the same backward as forward.
func IsPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	s = strings.ToLower(s)
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return false
		}
	}
	return true
}
