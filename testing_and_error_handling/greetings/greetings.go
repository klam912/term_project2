package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Features:
// 1. Error handling using log
// 2. Syntax of functions (specifying return values)

// Greet function accepts a string (first paranthesis) and returns a string and error (second paranthesis)
func Greet(name string) (string, error) {
	// Handle empty names
	if name == "" {
		return "", errors.New("empty name") // return the error message
	}

	message := fmt.Sprintf(randomMessage(), name)
	// message := fmt.Sprintf("Hello %v", name)

	return message, nil // return the message with no error
}

// GreetPeople greets multiple people with randomized messages (similar to Greet function)
// Note: to export your functions into another module, you must capitalize your function name
func GreetPeople(names []string) (map[string]string, error) {
	// Create a map (dict) to assign each name with a greeting message
	messages := make(map[string]string)

	// Loop through each name in names and assign a random greeting message
	for _, name := range names { // range based for loop (like enumerate in Python)
		message, error := Greet(name)
		if error != nil {
			return nil, error // return nothing and the error
		}
		// Assign the random message to messages
		messages[name] = message
	}
	return messages, nil
}

// randomMessage function that returns a randomized message
func randomMessage() string {
	// Use slice (Go's version of an array that can dynamically change size every time you add/remove items)
	greeting_messages := []string{
		"Hi %v!",
		"Salute %v!",
		"Howdy %v!",
	}

	// Get a randomized idx
	rand_idx := rand.Intn(len(greeting_messages))
	return greeting_messages[rand_idx]
}
