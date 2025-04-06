package main

import "fmt"

// A way for you to define your constraints (accepts either an int64 or float64)
type NumberConstraint interface {
	int64 | float64 // "|" is the union between the two datatypes
}

// Description: Generic function that multiplies numbers in a map
// The bracket specifies 2 generic datatypes (K and V)
// |=> K is of type comparable (pre-defined in Go) because a map's key must be type comparable
// |=> V follows the above-defined NumberConstraint, which accepts either int64 or float64
// The next pair of parantheses specifies the function's parameter called m, which is a map whose key is of type K and value is of type V
// Finally, the function returns a value of type V
func MultiplyNumbers[K comparable, V NumberConstraint](m map[K]V) V {
	var product V = 1       // initialize the product
	for _, val := range m { // loop through the map
		product *= val
	}

	return product
}

func main() {
	// Initialize a map for integer values
	map_of_ints := map[string]int64{
		"first":  55,
		"second": 245,
	}

	// Initialize a map for floating values
	map_of_floats := map[string]float64{
		"first":  34.5,
		"second": 776.44,
	}

	product_ints := MultiplyNumbers(map_of_ints)
	product_floats := MultiplyNumbers(map_of_floats)

	fmt.Println("Generic Products with Constraints:")
	fmt.Printf("Product of ints: %v \n", product_ints)
	fmt.Printf("Product of floats: %v \n", product_floats)

}
