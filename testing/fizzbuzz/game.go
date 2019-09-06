// Package fizzbuzz game calculator
package fizzbuzz

import "strconv"

const (
	divBy3     = "Fizz"
	divBy5     = "Buzz"
	divBy3And5 = divBy3 + divBy5
)

// Compute returns a FizzBuzz number based on a given input.
// Returns...
// `FizzBuzz if number div by 3 and 5.
// `Fizz if div by 3.
// `Buzz if div by 5
// Otherwise the given number
func compute(n int) (s string) {
	switch {
	case n%3 == 0 && n%5 == 0:
		return divBy3And5
	case n%3 == 0:
		return divBy3
	case n%5 == 0:
		return divBy5
	default:
		return strconv.Itoa(n)
	}
}
