package main

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

const (
	div3     = "Fizz"
	div5     = "Buzz"
	div3And5 = div3 + div5
)

// Error represents an error.
type Error string

// Error returns an error message.
func (e Error) Error() string {
	return string(e)
}

const (
	// ErrUnderRange represents an out of range error.
	ErrUnderRange = Error("number is under range (<=0)")
	// ErrOverRange represents an over range error.
	ErrOverRange = Error("number is over range (> 20)")
)

func main() {
	for i := 0; i <= 21; i++ {
		if r, err := play(i); err != nil {
			fmt.Printf("Err %+v\n", err)
		} else {
			fmt.Printf("%02d %v\n", i, r)
		}
	}
}

func play(n int) (string, error) {
	switch {
	case n <= 0:
		return "", errors.Wrapf(ErrUnderRange, "FizzBuzz with %d", n)
	case n > 20:
		return "", errors.Wrapf(ErrOverRange, "FizzBuzz with %d", n)
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
