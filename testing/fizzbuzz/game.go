// © 2018 Imhotep Software LLC. All rights reserved.

// Package fizzbuzz implements a FizzBuzz game.
package fizzbuzz

import "strconv"

const (
	divBy3 = "Fizz"
	// DivBy5 number divisible by 5
	divBy5 = "Buzz"
	// DivBy3And5 divisible by 3 and 5
	divBy3And5 = divBy3 + divBy5
)

// Compute a FizzBuzz number based on given input.
//
// 	Returns `FizzBuzz if number div by 3 and 5.
// 	`Fizz if div by 3.
// 	`Buzz if div by 5.
// 	`number otherwise.
func compute(n int) string {
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
