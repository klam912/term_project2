// Authors: Ken, Abby, Carlos
// Date: 2025-04-06
// Usage: go run generics.go
// Description: Demonstrates Go's generics function (allows a function to accept multiple datatypes)
// Source: https://go.dev/doc/tutorial/generics
package main

import (
	"fmt"
	"math/rand"
)

// A way for you to define your constraints (accepts either an int64 or float64)
type NumberConstraint interface {
	int | float64 // "|" is the union between the two datatypes
}

// Description: Function that does a mystery operation on any datatypes
// The bracket specifies 2 generic datatypes (K and V)
// |=> K is of type comparable (pre-defined in Go) because a map's key must be type comparable
// |=> V follows the above-defined NumberConstraint, which accepts either int64 or float64
// The next pair of parantheses specifies the function's parameter called m, which is a map whose key is of type K and value is of type V
// Finally, the function returns a value of type V
func MysteryOperator[K comparable, V NumberConstraint](m map[K]V) V {
	var result V = 1        // initialize the product
	for _, val := range m { // loop through the map
		result *= val + val
	}

	return result
}

func main() {
	// Initialize a map for integer values
	map_of_ints := map[string]int{
		"first":  rand.Intn(101),
		"second": rand.Intn(101),
	}

	// Initialize a map for floating values
	map_of_floats := map[string]float64{
		"first":  rand.Float64(),
		"second": rand.Float64(),
	}

	result_ints := MysteryOperator(map_of_ints)
	result_floats := MysteryOperator(map_of_floats)

	fmt.Println("Generic Products with Constraints:")
	fmt.Printf("Product of ints: %v \n", result_ints)
	fmt.Printf("Product of floats: %v \n", result_floats)

}
