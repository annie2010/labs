package greeter

import (
	"net/http"
)

// Response greeting message format.
type Response struct {
	Greeting, Name string
	Age            int
}

// SetupRoutes for various endpoints.
func SetupRoutes() {
	// YOUR CODE ...
}

// GreetingHandler greats incoming users.
func GreetingHandler(w http.ResponseWriter, r *http.Request) {
	// YOUR CODE...
}

// NoMatchHandler is used when no known endpoints.
func NoMatchHandler(w http.ResponseWriter, r *http.Request) {
	// YOUR CODE...
}
