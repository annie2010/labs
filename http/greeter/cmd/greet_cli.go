package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/imhotepio/letsgo_labs/http/greeter"
)

func main() {
	var (
		name = flag.String("name", "no one", "A user's name")
		age  = flag.Int("age", 42, "A user's age")
		url  = flag.String("url", "localhost:4500", "The greeting service url")
	)
	flag.Parse()

	params := map[string]string{
		"name": *name,
		"age":  strconv.Itoa(*age),
	}
	r, err := callGreet(*url, params)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("[DialAGreeter] says `%s\n", r.Greeting)
}

func callGreet(url string, params map[string]string) (greeter.Response, error) {
	return get(url, params)
}

func get(host string, params map[string]string) (greeter.Response, error) {
	var resp greeter.Response

	u, err := url.Parse(fmt.Sprintf("http://%s", host))
	if err != nil {
		return resp, err
	}
	u.Path = "/greet"

	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	rr, err := http.Get(u.String())
	if err != nil {
		return resp, err
	}
	defer rr.Body.Close()

	if rr.StatusCode != 200 {
		return resp, fmt.Errorf("Response not ok %d", rr.StatusCode)
	}

	err = json.NewDecoder(rr.Body).Decode(&resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
