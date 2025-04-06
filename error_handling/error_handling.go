// Authors: Ken, Abby, Carlos
// Date: 2025-04-06
// Usage: main
// Description: Demonstrates Go's error handling
package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// Get user input
// Note: for a function to be imported to another module, must capitalize function name
func GetInput() (string, error) { // first parantheses has no input and second parantheses return a string and error
	reader := bufio.NewReader(os.Stdin)        // creates a buffered reader that reads from standard input
	fmt.Print("Enter text: ")                  // Prompts user input
	text, _ := reader.ReadString('\n')         // reads user input until a newline is entered
	text = strings.Replace(text, "\n", "", -1) // removes newline character from the end of input string

	// If an empty text, return the error message
	if text == "" {
		return "", errors.New("empty input")
	}

	return text, nil // if successful, return the input and nil (no errors)
}

func main() {
	// Customize a Logger
	log.SetPrefix("Error: ") // customizes the prefix of the error output
	log.SetFlags(0)          // disables unecessary information (time, number line, etc. of the error)

	// Error handling
	message, error := GetInput()

	if error != nil {
		log.Fatal(error) // prints the error message to the console and exit the program
	}

	fmt.Println(message)

}
