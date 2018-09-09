package main

import (
	"fmt"
	"strconv"
)

const (
	div3     = "Fizz"
	div5     = "Buzz"
	div3And5 = div3 + div5
)

// YOUR CODE define your sentinel errors
var ()

func main() {
	for i := 0; i <= 21; i++ {
		fmt.Printf("%02d ", i)
		if r, err := play(i); err != nil {
			fmt.Printf("%v\n", err)
		} else {
			fmt.Printf("%v\n", r)
		}
	}
}

func play(n int) (string, error) {
	var s string

	// YOUR CODE... Check for number out of range

	switch {
	case n%3 == 0 && n%5 == 0:
		s = div3And5
	case n%3 == 0:
		s = div3
	case n%5 == 0:
		s = div5
	default:
		s = strconv.Itoa(n)
	}

	return s, nil
}
