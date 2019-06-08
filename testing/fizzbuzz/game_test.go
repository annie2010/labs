// © 2019 Imhotep Software LLC. All rights reserved.

package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompute(t *testing.T) {
	uu := map[int]string{
		0:  divBy3And5,
		1:  "1",
		3:  divBy3,
		4:  "4",
		5:  divBy5,
		15: divBy3And5,
	}

	for k, v := range uu {
		assert.Equal(t, v, compute(k))
	}
}
