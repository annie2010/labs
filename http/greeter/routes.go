package greeter

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Response greeting message format.
type Response struct {
	Greeting, Name string
	Age            int
}

// SetupRoutes for various endpoints.
func SetupRoutes() {
	// Setup handlers for / and /greet
	// YOUR CODE ...
}

// GreetingHandler greats incoming users.
func GreetingHandler(w http.ResponseWriter, r *http.Request) {
	// Pull the name and age out of the request and
	// call the greeter api to produce a greeting
	// YOUR CODE...

	resp := Response{
		Greeting: msg,
		Name:     name,
		Age:      age,
	}
	buff, err := json.Marshal(&resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, string(buff))
}

// NoMatchHandler is used when no known endpoints.
func NoMatchHandler(w http.ResponseWriter, r *http.Request) {
	// Should produce an http error `no matching routes` with a 400 status code.
	// YOUR CODE...
}
