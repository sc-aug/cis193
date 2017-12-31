// Homework 8: CLI and Regex
// Due April 4, 2017 at 11:59pm
package main

import (
	"fmt"
	"regexp"
	"unicode"
	"flag"
)

// Problem 1: CLI
// Write a command line interface that prints out sequences of numbers.
//
// Usage of hw8:
// 	hw8 [flags] # prints out the sequence of numbers, each on a new line
// Flags:
//   -start int
//     	starting integer for the sequence (default 0)
//   -end   int
//      ending integer for the sequence, not inclusive (default 0)
//   -step  int
//      amount to skip by in each iteration (default 1)
//
// For example, executing `./hw8 -start=2 -end=5` should print out:
// 2
// 3
// 4
//
// Executing `./hw8 -start=2 -end=7 -step=3` should print out:
// 2
// 5
//
// Executing `./hw8 -start=10 -end=7 -step=-1` should print out:
// 10
// 9
// 8
//
// If the parameters are invalid (eg: positive step and start > end or
// negative step and start < end or invalid parameter values passed in),
// print out an error message using `log.Print(ln|f)?`.
//
// Feel free to do this section directly in the main() function.

func main() {
  s := flag.Int("start", 0, "an int")
	e := flag.Int("end", 0, "an int")
	p := flag.Int("step", 0, "an int")
	flag.Parse()
	for i:= *s; (*e-i) * *p > 0; i += *p {
		fmt.Println(i)
	}
}


// GetEmails takes in string input and returns a string slice of the
// emails found in the input string.
//
// Use regexp to extract all of the emails from the input string.
// Each email consists of the email name + "@" + domain + "." + top level domain.
// The email name should consist of only letters, numbers, underscores and dots.
// The domain should consist of only letters or dots.
// The top level domain must be "com", "org", "net" or "edu".
// between the domain and tld.
//
// You can assume that all email addresses will be surrounded by whitespace.
func GetEmails(s string) []string {
	reg := regexp.MustCompile(`[A-Za-z0-9._+-]+@[A-Za-z.]+.(com|net|edu|org)`)
	emails := reg.FindAllString(s, -1)
	return emails
}

// GetPhoneNumbers takes in string input and returns a string slice of the
// phone numbers found in the input string.
//
// Use regexp to extract all of the phone numbers from the input string.
// Here are the formats phone numbers can be in for this problem:
// 215-555-3232
// (215)-555-3232
// 215.555.3232
// 2155553232
// 215 555 3232
//
// For your output, you should return a string slice of phone numbers with
// just the numbers (eg: "2158887744")
//
// You can assume that all phone numbers will be surrounded by whitespace.
func GetPhoneNumbers(s string) []string {
	reg := regexp.MustCompile(`[\\(\\)0-9\s.+-]{10,}`)
	numstr := reg.FindAllString(s, -1)
	nums := make([]string, len(numstr))
	for i, ns := range numstr {
		nums[i] = ParsePhone(ns)
	}
	return nums
}

// from hw1
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
	return string(num)
}
