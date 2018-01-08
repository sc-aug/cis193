package main

import (
	"fmt"
	"testing"
)

func TestScrapeHackerNews(t *testing.T) {
	ScrapeHackerNews(10)
}

func TestGetEmails(t *testing.T) {
	fmt.Println(GetEmails())
}

func TestGetCountryGDP(t *testing.T) {
	// GetCountryGDP("Lithuania")Colombia
	fmt.Println(GetCountryGDP("Colombia"))
}
