package greeter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreeter1(t *testing.T) {
	assert.Equal(t, "Hello, Fernand! You are 42 old today...", greeter1("Fernand", 42))
}

func TestGreeter2(t *testing.T) {
	// YOUR_CODE
}

func TestGreeter3(t *testing.T) {
	// YOUR_CODE
}

func BenchmarkGreeter1(b *testing.B) {
	// YOUR_CODE
}

func BenchmarkGreeter2(b *testing.B) {
	// YOUR_CODE
}

func BenchmarkGreeter3(b *testing.B) {
	// YOUR_CODE
}
