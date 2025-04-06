// Authors: Ken, Abby, Carlos
// Date: 2025-04-06
// Usage: go run defer.go
// Description: Demonstrates that functions can be passed as arguments and returned from other functions

package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func minus(x int, y int) int {
	return x - y
}

func higher_order_function(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

func main() {
	// Assign functions as variables
	addFunction := add
	minusFunction := minus

	// Pass function variables as arguments
	result1 := higher_order_function(34, 56, addFunction)
	result2 := higher_order_function(356, 234, minusFunction)

	// Print results
	fmt.Println("Result 1: ", result1)
	fmt.Println("Result 2: ", result2)
}
