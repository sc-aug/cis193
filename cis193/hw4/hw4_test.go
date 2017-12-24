package main

import (
	"log"
	"os"
	"testing"
)

func TestFileSum(t *testing.T) {
	file_in := "file_sum.txt"
	file_out := "out.txt"
	FileSum(file_in, file_out)
}

func TestIOSum(t *testing.T) {
	// filename
	file_in := "file_sum.txt"
	file_out := "out.txt"
	// Open input file
	fi, err := os.Open(file_in)
	defer fi.Close()
	if err != nil {
		log.Fatal(err)
	}
	// Open output file
	fo, err := os.OpenFile(file_out, os.O_WRONLY|os.O_CREATE, 0644)
	defer fo.Close()
	if err != nil {
		log.Fatal(err)
	}

	IOSum(fi, fo)
}
