// Authors: Ken, Carlos, Abby
// Date: 2025-04-06
// Usage: main

package main

import (
	"fmt"
	"log"
	"term_project/greetings"
)

// Features:
// 1. Error handling using log
// 2. Multi-module

// Main function
func main() {
	// Set the properties of the logger
	log.SetPrefix("Error: ") // sets the prefix of the error output
	log.SetFlags(0)          // disable meta-data printing (number line, time, source file)

	// Define name
	a_name := "Ken"

	// Greet a_name
	message, error := greetings.Greet(a_name)

	// Define list of names
	name_list := []string{
		"Carlos",
		"Ken",
		"Abby",
	}

	// Greet everyone
	messages, error := greetings.GreetPeople(name_list)

	// Handle the error
	if error != nil { // if an error is returned, then print it to the console
		log.Fatal(error)
	}

	// Print message for a_name
	fmt.Println("Greeting one person: ")
	fmt.Println(message)

	// Print message for name_list
	fmt.Println("Greeting multiple people: ")
	for _, val := range messages {
		fmt.Println(val)
	}
}
