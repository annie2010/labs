package greeter

import (
	"errors"
	"fmt"
)

func greet(name string, age int) (string, error) {
	if len(name) == 0 {
		return "", errors.New("you must provide a name")
	}
	return fmt.Sprintf("Greetings %s! You are now %d years old...", name, age), nil
}
