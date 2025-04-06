package greetings

import (
	"regexp"
	"testing"
)

// Feature:
// 1. Built-in testing

// Check for a valid return value
// Start with Test and Name is customizable (e.g., TestSomething, TestName, TestHaha)
func TestName(t *testing.T) {
	name := "Ken"
	want := regexp.MustCompile(`\b` + name + `\b`) // valid if name is surrounded by non-word characters
	message, error := Greet("Ken")                 // run the function

	// if didn't match, then print error message
	if !want.MatchString(message) || error != nil {
		// Standard error print format
		// %q puts message in quoted string
		// %v puts error as a value
		// %#q puts want as quoted string in Go's syntax literal (with escape characters)
		t.Errorf(`Greet("Ken") = %q, %v, want match for %#q, nil`, message, error, want)
	}
}
