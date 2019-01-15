package main

import (
	"io"
	"net/http"
)

type hotdog int
// implements type Handler
func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

type hotcat int
// implements type Handler
func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	var d hotdog
	var c hotcat

	mux := http.NewServeMux()
	mux.Handle("/dog/", d)
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux)
}
