// © 2019 Imhotep Software LLC. All rights reserved.

// Package fizzbuzz implements a FizzBuzz game.
package fizzbuzz

import "strconv"

const (
	divBy3     = "Fizz"
	divBy5     = "Buzz"
	divBy3And5 = divBy3 + divBy5
)

// YOUR_CODE
func Compute(n int) string {
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
