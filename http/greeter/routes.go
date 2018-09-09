package greeter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Response greeting message format.
type Response struct {
	Greeting, Name string
	Age            int
}

// SetupRoutes for various endpoints.
func SetupRoutes() {
	http.HandleFunc("/greet", GreetingHandler)
	http.HandleFunc("/", NoMatchHandler)
}

// GreetingHandler greats incoming users.
func GreetingHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	age, err := strconv.Atoi(r.URL.Query().Get("age"))
	if err != nil {
		http.Error(w, "invalid age specified", http.StatusExpectationFailed)
		return
	}

	msg, err := greet(name, age)
	if err != nil {
		http.Error(w, "you must provide a user name", http.StatusExpectationFailed)
		return
	}

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
	http.Error(w, fmt.Sprintf("no matching routes for %s", r.URL), http.StatusBadRequest)
}
