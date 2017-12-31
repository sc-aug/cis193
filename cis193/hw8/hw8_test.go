package main

import (
	"fmt"
	"testing"
)

func TestGetEmails(t *testing.T) {
	text := `alur@cis.upenn.edu ashivani@seas.upenn.edu badler@seas.upenn.edu bhusnur4@seas.upenn.edu blaze@cis.upenn.edu ccb@cis.upenn.edu pjcozzi@gmail.com kostas@cis.upenn.edu susan@cis.upenn.edu devietti@cis.upenn.edu eeaton@seas.upenn.edu tfarmer@seas.upenn.edu jean@cis.upenn.edu rajivg@camden.rutgers.edu sudipto@seas.upenn.edu ahae@cis.upenn.edu fbrett@cis.upenn.edu nadiah@cis.upenn.edu zives@cis.upenn.edu cffjiang@seas.upenn.edu kannan@cis.upenn.edu mkearns@cis.upenn.edu sanjeev@cis.upenn.edu shlane@cis.upenn.edu lee@cis.upenn.edu liuv@cis.upenn.edu boonloo@cis.upenn.edu amally@seas.upenn.edu mitch@cis.upenn.edu paulmcb@cis.upenn.edu mintz@cis.upenn.edu cdmurphy@cis.upenn.edu mhnaik@cis.upenn.edu nenkova@seas.upenn.edu linhphan@cis.upenn.edu bcpierce@cis.upenn.edu aaroth@cis.upenn.edu danroth@cis.upenn.edu swapneel@cis.upenn.edu jshi@cis.upenn.edu jms@cis.upenn.edu sokolsky@cis.upenn.edu val@cis.upenn.edu cjtaylor@central.cis.upenn.edu ungar@cis.upenn.edu weimerj@seas.upenn.edu sweirich@cis.upenn.edu stevez@cis.upenn.edu`
	for _, e := range GetEmails(text) {
		fmt.Println(e)
	}
}

func TestGetPhoneNumbers(t *testing.T) {
	text := `// GetPhoneNumbers takes in string input and returns a string slice of the
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
	// just the numbers (eg: "2158887744")`
	for i, p := range GetPhoneNumbers(text) {
		fmt.Println(i, p)
	}
}
