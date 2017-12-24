// Homework 7: Web Scraping
// Due March 28, 2017 at 11:59pm
package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("http://uszip.com/zip/19104")
	if err != nil {
		log.Fatal(err)
	}
	city := doc.Find("div.zip-data hgroup h2 strong").Eq(0).Text()
	fmt.Println(city)
}

// News is a Hacker News article listing
type News struct {
	Points   int
	Title    string
	Username string
	URL      string
}

// NewsSlice is a slice of News pointers
type NewsSlice []*News

// ScrapeHackerNews scrapes the website "https://news.ycombinator.com/" using goquery and returns
// information on the first n posts.
//
// For each post, the attributes to be extracted are: points, title, username and url.
// This data should be returned as a NewsSlice, where NewsSlice is a custom slice of News structs.
//
// For example, for the sample image located at `https://www.cis.upenn.edu/~cis193/homeworks/hn.png`,
// the struct would look like:
// News{24, "QEMU(TCG): user-to-root privesc inside VM via bad translation caching",
// "webaholic", "https://bugs.chromium.org/p/project-zero/issues/detail?id=1122"}.
//
// If n is greater than the number of total posts available (which should be 30), return data from
// the all of the available posts (all thirty).
func ScrapeHackerNews(n int) NewsSlice {
	// TODO
	return nil
}

// GetEmails returns a string slice of the emails found on the given URL.
//
// Scenario: you are a student enthusiastic about spreading awareness about Go. To effectively
// market Go, you decide to email Penn CIS professors about the wonders of the Go programming
// language. In this function, use goquery to extract the email addresses from the URL
// "http://www.cis.upenn.edu/about-people/" and return them as a string slice. This will involve you
// having to investigate where and how emails are located on the webpage.
// Note: you should have 47 total emails returned.
func GetEmails() []string {
	// TODO
	return nil
}

// CountryData has GDP information on a country
type CountryData struct {
	Country string
	GDP     string
}

// GetCountryGDP takes in a string country name and returns the GDP (in millions) as
// an integer. Information on the country is found by concurrently scraping a hidden website with
// data on countries scattered on many pages.
//
// Scenario: imagine you are a spy and you have discovered a URL with top secret GDP information:
// "https://www.cis.upenn.edu/~cis193/scraping/9828772efc2bd314a277c8880695dea2.html". This webpage
// has a country name and the GDP (in millions of US Dollars). It also has links to two other
// country's webpages. Based on intelligence you've received, every country has a webpage on this
// website with information about it, but you do not know the URL for each page. You can assume that
// none of the page links lead you to a cycle and every country can be reached from a path from the
// initial URL that you are given. So, for this function, you will need to traverse from the initial
// url to every webpage link you encounter in order to find information on the target `country`
// string. Since time is of the essence, you want to use concurrency to scrape webpages
// simultaneously. Note that for this function, we only care about getting the GDP for the input
// `country` string. You may find it useful to use the CountryData struct to send country
// information between goroutines.
//
// To prevent the function from getting stuck if an invalid `country` string is entered,
// you should also implement a timeout that will automatically return an error after 10 seconds
// if the program hasn't already finished terminating.
//
// Feel free to make and use helper functions for this function. To help with testing this
// function, we know from intelligence reports that the GDP for "Canada" is 1532343 and
// the GDP for "Colombia" is 274135.
func GetCountryGDP(country string) (int, error) {
	// TODO
	return 0, nil
}
