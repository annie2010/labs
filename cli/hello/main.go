package main

import (
	"fmt"
)

func greet(name string, age int) string {
	return fmt.Sprintf("Hello, my name is %s and I am %d years old!", name, age)
}

func main() {
	// YOUR_CODE...

	fmt.Println(greet("Blee", 10))
}
