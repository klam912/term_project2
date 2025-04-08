// Authors: Ken, Abby, Carlos
// Date: 2025-04-06
// Usage: go test -fuzz=FuzzMSE
// Description: Demonstrates Go's unique testing function (Fuzz Testing)

// What is Fuzz testing:
// Fuzzing is a type of automated testing which continuously manipulates inputs to a program to find bugs.
// It can reach edge cases which humans would often miss and allows you to find security exploits and vulnerabilities.
// Source: https://go.dev/doc/security/fuzz/

package main

import (
	"errors"
	"testing"
)

// Create a custom mean squared error function
func MSE(y_pred []float64, y_true []float64) (float64, error) {
	// Handles when length of two slices are different
	if len(y_pred) != len(y_true) {
		return -1, errors.New("unequal length between y_pred and y_true")
	}

	// Formula: 1/n (sum(y_pred - y_true)^2)
	var sum_squared_differences float64 = 0.0
	for i := 0; i < len(y_pred); i++ {
		difference := y_pred[i] - y_true[i]
		squared_difference := difference * difference
		sum_squared_differences += squared_difference
	}

	// mse := sum_squared_differences / float64(len(y_pred)) * -1 // uncomment this to break the mse function and see the FuzzTest return error
	mse := sum_squared_differences / float64(len(y_pred))

	return float64(mse), nil

}

// Function: y = 3x + 5
func true_function(x []float64) []float64 {
	y_pred := make([]float64, len(x)) // empty slice (dynamic version of an array (C++'s vector))
	for i := 0; i < len(x); i++ {
		y_pred[i] = 3*x[i] + 5
	}

	return y_pred
}

// Model
func model(x []float64) []float64 {
	y_pred := make([]float64, len(x)) // empty slice (dynamic version of an array (C++'s vector))
	for i := 0; i < len(x); i++ {
		y_pred[i] = 2*x[i] + 5
	}

	return y_pred
}

// Generate x values given a multiplier
func generateXValues(multiplier float64) []float64 {
	x := []float64{1.0, 2.0, 3.0, 4.0, 5.0} // Start with a small set of values
	for i := range x {
		x[i] = x[i] * float64(multiplier) // Multiply each value by the multiplier
	}
	return x
}

// Fuzz Test
func FuzzMSE(f *testing.F) {
	// Add initial inputs for the fuzz testing
	f.Add(10.0)
	f.Add(100.0)
	f.Add(-10.0)
	f.Add(0.00000004)

	// Fuzz testing with an x_multiplier
	f.Fuzz(func(t *testing.T, x_multiplier float64) {
		// Generate x values
		x := generateXValues(x_multiplier)

		y_pred := model(x)
		y_true := true_function(x)

		// Call the MSE function
		mse, error := MSE(y_pred, y_true)
		if error != nil {
			t.Errorf("Unexpected error!")
		} else {
			if mse < 0 {
				t.Errorf("MSE shouldn't be negative: mse = %v", mse)
			}
		}
	})
}
