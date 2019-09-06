// © 2019 Imhotep Software LLC. All rights reserved.

package fizzbuzz_test

import (
	"testing"

	"github.com/gopherland/labs/documentation/fizzbuzz"
	"github.com/stretchr/testify/assert"
)

func TestCompute(t *testing.T) {
	uu := map[int]string{
		0:  "FizzBuzz",
		1:  "1",
		3:  "Fizz",
		4:  "4",
		5:  "Buzz",
		15: "FizzBuzz",
	}

	for k, v := range uu {
		assert.Equal(t, v, fizzbuzz.Compute(k))
	}
}

// YOUR_CODE
func ExampleCompute_DivisibleBy3And5() {
	// YOUR_CODE
}

// YOUR_CODE
func ExampleCompute_DivisibleBy3() {
	// YOUR_CODE
}

// YOUR_CODE
func ExampleCompute_DivisibleBy5() {
	// YOUR_CODE
}

// YOUR_CODE
func ExampleCompute_NoFizzBuzz() {
	// YOUR_CODE
}
