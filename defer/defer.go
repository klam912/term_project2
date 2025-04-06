// Authors: Ken, Abby, Carlos
// Date: 2025-04-06
// Usage: go run defer.go
// Description: Demonstrates a defer statement, which defers the execution of a function until the surrounding function returns

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFile(filename string) {
	// Open the file
	file, error := os.Open(filename)
	if error != nil {
		log.Fatalf("Failure to open file: %v\n", error)
	}

	// Defer these line to ensure that the file closes when we're done even if the function returns errors early
	defer fmt.Println("Closing file")
	defer file.Close()

	fmt.Printf("Reading file: %v\n", filename)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if error := scanner.Err(); error != nil {
		log.Fatalf("Error reading file: %v\n", error)
	}
}

func main() {
	filename := "example.txt"
	readFile(filename)
	fmt.Println("Finished reading")
}
