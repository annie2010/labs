package main

import (
	"log"
	"net/http"
)

const httpPort = ":4500"

func main() {
	// YOUR CODE...

	log.Printf("Dial a Greeter is listening [%s]", httpPort)
	log.Fatal(http.ListenAndServe(httpPort, nil))
}
