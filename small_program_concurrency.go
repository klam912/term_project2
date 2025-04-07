// Author: Carlos Moncada Soto, Abby Baron, Ken Lam
// Program Name: small_program_concurrency.go
// Date: 04/07/2025 
// Description: A small program which demonstrates concurrency using goroutines.
//              It simulates downloading files concurrently.
// Usage: go run small_program_concurrency.go

// References: https://www.golang-book.com/books/intro/10

// The 'main' package is the starting point of any Go program.
package main

import (
    "fmt" // 'fmt' is used for printing to the console.
    "time" // 'time' allows us to simulate delays
)

// downloadFile simulates downloading a file by printing a start message,
// pausing the program for 2 seconds (to represent download time),
// and then printing a completion message.
func downloadFile(fileName string) {
    fmt.Printf("Starting download of %s...\n", fileName)

    time.Sleep(2 * time.Second)

    fmt.Printf("Finished downloading %s!\n", fileName)
}

// The main function is the entry point of the program and runs automatically.
func main() {
    fmt.Println("beginning the program...")

    // Each of these lines starts a new goroutine using the 'go' keyword.
    // A goroutine is a lightweight thread managed by Go for concurrent tasks.
    go downloadFile("file1.txt")
    go downloadFile("file2.txt")
    go downloadFile("file3.txt")

    // This line runs immediately after the goroutines start.
    // Because goroutines run in the background
    fmt.Println("The main function is still running here and the program can do other things, will not just wait")

    // Pause the main function for 3 seconds so the downloads have time to finish.
    time.Sleep(3 * time.Second)

    fmt.Println("All downloads started! Main function is done.")
}