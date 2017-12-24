// Homework 4: Concurrency
// Due February 21, 2017 at 11:59pm
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	// Feel free to use the main function for testing your functions
	hello := map[string]string{
		"こんにちは": "世界",
		"你好":    "世界",
		"안녕하세요": "세계",
	}
	for k, v := range hello {
		fmt.Printf("%s, %s\n", strings.Title(k), v)
	}
}

// Problem 1a: File processing
// You will be provided an input file consisting of integers, one on each line.
// Your task is to read the input file, sum all the integers, and write the
// result to a separate file.

// FileSum sums the integers in input and writes them to an output file.
// The two parameters, input and output, are the filenames of those files.
// You should expect your input to end with a newline, and the output should
// have a newline after the result.
func FileSum(input, output string) {
	// Open input file
	fi, err := os.Open(input)
	defer fi.Close()
	if err != nil {
		log.Fatal(err)
	}
	// Open output file
	// fo, err := os.Open(output)
	fo, err := os.OpenFile(output, os.O_WRONLY|os.O_CREATE, 0644)
	defer fo.Close()
	if err != nil {
		log.Fatal(err)
	}
	// SUM
	sum := 0
	// Read buffer
	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("strconv.Atoi() error: %v\n", err)
		}
		sum += num
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("scanner error:", err)
	}
	// Write buffer
	writer := bufio.NewWriter(fo)
	writer.WriteString(strconv.Itoa(sum))
	writer.Flush()
}

// Problem 1b: IO processing with interfaces
// You must do the exact same task as above, but instead of being passed 2
// filenames, you are passed 2 interfaces: io.Reader and io.Writer.
// See https://golang.org/pkg/io/ for information about these two interfaces.
// Note that os.Open returns an io.Reader, and os.Create returns an io.Writer.

// IOSum sums the integers in input and writes them to output
// The two parameters, input and output, are interfaces for io.Reader and
// io.Writer. The type signatures for these interfaces is in the Go
// documentation.
// You should expect your input to end with a newline, and the output should
// have a newline after the result.
func IOSum(input io.Reader, output io.Writer) {
	// SUM
	sum := 0
	// Read buffer
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("strconv.Atoi() error: %v\n", err)
		}
		sum += num
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("scanner error:", err)
	}
	// Write buffer
	writer := bufio.NewWriter(output)
	writer.WriteString(strconv.Itoa(sum))
	writer.Flush()
}

// Problem 2: Concurrent map access
// Maps in Go [are not safe for concurrent use](https://golang.org/doc/faq#atomic_maps).
// For this assignment, you will be building a custom map type that allows for
// concurrent access to the map using mutexes.
// The map is expected to have concurrent readers but only 1 writer can have
// access to the map.

// PennDirectory is a mapping from PennID number to PennKey (12345678 -> adelq).
// You may only add *private* fields to this struct.
// Hint: Use an embedded sync.RWMutex, see lecture 2 for a review on embedding
type PennDirectory struct {
	mu        sync.RWMutex
	directory map[int]string
}

// Add inserts a new student to the Penn Directory.
// Add should obtain a write lock, and should not allow any concurrent reads or
// writes to the map.
// You may NOT write over existing data - simply raise a warning.
func (d *PennDirectory) Add(id int, name string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if n, ok := d.directory[id]; ok {
		log.Printf("Warning: entry existed id:%d name:\n", id, n)
	} else {
		d.directory[id] = name
	}
}

// Get fetches a student from the Penn Directory by their PennID.
// Get should obtain a read lock, and should allow concurrent read access but
// not write access.
func (d *PennDirectory) Get(id int) string {
	d.mu.RLock()
	defer d.mu.RUnlock()
	name, ok := d.directory[id]
	if !ok {
		log.Printf("Warning: no such entry existed")
	}
	return name
}

// Remove deletes a student to the Penn Directory.
// Remove should obtain a write lock, and should not allow any concurrent reads
// or writes to the map.
func (d *PennDirectory) Remove(id int) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if _, ok := d.directory[id]; !ok {
		log.Printf("Warning: no such entry existed")
	} else {
		delete(d.directory, id)
	}
}
