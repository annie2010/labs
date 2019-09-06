// © 2019 Imhotep Software LLC. All rights reserved.

package fizzbuzz_test

import (
	"fmt"
	"testing"

	"github.com/gopherland/labs/documentation/fizzbuzz"
	"github.com/stretchr/testify/assert"
)

var Output string

func TestCompute(t *testing.T) {
	uu := map[int]string{
		0:  fizzbuzz.DivBy3And5,
		1:  "1",
		3:  fizzbuzz.DivBy3,
		4:  "4",
		5:  fizzbuzz.DivBy5,
		15: fizzbuzz.DivBy3And5,
	}

	for k, v := range uu {
		assert.Equal(t, v, fizzbuzz.Compute(k))
	}
}

func BenchmarkCompute(b *testing.B) {
	var out string
	for i := 0; i < b.N; i++ {
		out = fizzbuzz.Compute(i)
	}
	Output = out
}

// Returns `FizzBuzz if number is div by 3 and 5.
func ExampleCompute_DivisibleBy3And5() {
	fmt.Println(fizzbuzz.Compute(15))
	// Output:
	// FizzBuzz
}

// Returns `Fizz if number is div by 3.
func ExampleCompute_DivisibleBy3() {
	fmt.Println(fizzbuzz.Compute(3))
	// Output:
	// Fizz
}

// Returns `Buzz if number is div by 5.
func ExampleCompute_DivisibleBy5() {
	fmt.Println(fizzbuzz.Compute(5))
	// Output:
	// Buzz
}

// Returns number if not a Fizzbuzz number.
func ExampleCompute_NotFizzBuzz() {
	fmt.Println(fizzbuzz.Compute(2))
	// Output:
	// 10
}
