// © 2019 Imhotep Software LLC. All rights reserved.

// Package fizzbuzz implements a FizzBuzz game.
package fizzbuzz

import "strconv"

const (
	// DivBy3 number divisible by 3
	DivBy3 = "Fizz"
	// DivBy5 number divisible by 5
	DivBy5 = "Buzz"
	// DivBy3And5 divisible by 3 and 5
	DivBy3And5 = DivBy3 + DivBy5
)

// Compute a FizzBuzz number based on given input.
//
// 	Returns
//	`FizzBuzz if number is divisible by 3 and 5
// 	`Fizz if number is divisible by 3
// 	`Buzz if number is divisible by 5
// 	`number otherwise.
func Compute(n int) string {
	switch {
	case n%3 == 0 && n%5 == 0:
		return DivBy3And5
	case n%3 == 0:
		return DivBy3
	case n%5 == 0:
		return DivBy5
	default:
		return strconv.Itoa(n)
	}
}
