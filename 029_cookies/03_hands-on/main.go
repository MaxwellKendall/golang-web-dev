package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/increment", increment)
	http.ListenAndServe(":8080", nil)
}

func index(rw http.ResponseWriter, req *http.Request) {
	http.SetCookie(rw, &http.Cookie{Name: "count", Value: "0"})
}

func increment(rw http.ResponseWriter, req *http.Request) {
	counter, err := req.Cookie("count")

	if err != nil {
		fmt.Fprint(rw, err)
	}

	count, err := strconv.Atoi(counter.Value)

	if err != nil {
		fmt.Fprint(rw, err)
	}

	count++

	http.SetCookie(rw, generateCounterCookie(count))

}

func generateCounterCookie(count int) *http.Cookie {
	return &http.Cookie{
		Name:  "count",
		Value: strconv.Itoa(count),
	}
}
