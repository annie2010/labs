package main

import (
	"log"
	"net/http"

	"github.com/gopherland/labs/http/greeter"
)

const httpPort = ":4500"

func main() {
	greeter.SetupRoutes()

	log.Printf("Dial a Greeter is listening [%s]", httpPort)
	log.Fatal(http.ListenAndServe(httpPort, nil))
}
