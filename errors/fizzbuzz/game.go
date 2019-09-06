package main

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	div3     = "Fizz"
	div5     = "Buzz"
	div3And5 = div3 + div5
)

type (
	// ErrUnderRange represents a number under range error.
	ErrUnderRange struct{}
	// ErrOverRange represents a number over range error.
	ErrOverRange struct{}
)

func (e ErrUnderRange) Error() string {
	return "number is under range (<=0)"
}

func (e ErrOverRange) Error() string {
	return "number is over range (>20)"
}

func main() {
	for i := 0; i <= 21; i++ {
		r, err := play(i)
		if err == nil {
			fmt.Printf("%02d %v\n", i, r)
			continue
		}

		// Check err under range if so print unwrapped error
		var e ErrUnderRange
		if errors.As(err, &e) {
			fmt.Printf("Boom!! %s\n", e)
			continue
		}

		// Otherwise print wrapped error
		fmt.Println("Bam!!", err)
	}
}

func play(n int) (string, error) {
	switch {
	case n <= 0:
		return "", fmt.Errorf("Invalid FizzBuzz# (%d) -- %w", n, ErrUnderRange{})
	case n > 20:
		return "", fmt.Errorf("Invalid FizzBuzz# (%d) -- %w", n, ErrOverRange{})
	case n%3 == 0 && n%5 == 0:
		return div3And5, nil
	case n%3 == 0:
		return div3, nil
	case n%5 == 0:
		return div5, nil
	default:
		return strconv.Itoa(n), nil
	}
}
