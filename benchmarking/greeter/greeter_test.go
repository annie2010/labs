package greeter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var Output string

const expected = "Hello, Fernand! You are 42 old today..."

func TestGreeter1(t *testing.T) {
	assert.Equal(t, expected, greeter1("Fernand", 42))
}

func TestGreeter2(t *testing.T) {
	assert.Equal(t, expected, greeter2("Fernand", 42))
}

func TestGreeter3(t *testing.T) {
	assert.Equal(t, expected, greeter3("Fernand", 42))
}

func TestGreeter4(t *testing.T) {
	assert.Equal(t, expected, greeter4("Fernand", 42))
}

func BenchmarkGreeter1(b *testing.B) {
	n, a := "Fernand", 30
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		greeter1(n, a)
	}
}

func BenchmarkGreeter2(b *testing.B) {
	n, a := "Fernand", 30
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		greeter2(n, a)
	}
}

func BenchmarkGreeter3(b *testing.B) {
	n, a := "Fernand", 30
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		greeter3(n, a)
	}
}

func BenchmarkGreeter4(b *testing.B) {
	n, a := "Fernand", 30
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		greeter4(n, a)
	}
}
