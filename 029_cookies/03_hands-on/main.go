package main

import (
	"io"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/increment", increment)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	// http.SetCookie(res, &http.Cookie{Name: "count", Value: "0"})
	io.WriteString(res, "Nicely done...?")
}

func increment(res http.ResponseWriter, req *http.Request) {
	counter, err := req.Cookie("count")

	if err == http.ErrNoCookie {
		counter = generateCounterCookie(0)
	} else if err != nil {
		fmt.Fprint(res, err)
	}

	count, err := strconv.Atoi(counter.Value)

	if err != nil {
		fmt.Fprint(res, err)
	}

	count++

	http.SetCookie(res, generateCounterCookie(count))
}

func generateCounterCookie(count int) *http.Cookie {
	return &http.Cookie{
		Name:  "count",
		Value: strconv.Itoa(count),
	}
}
