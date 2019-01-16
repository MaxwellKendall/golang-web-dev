package main

import (
	"io"
	"net/http"
)

func indexHandler(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "yo INDEX wutup dis is max")	
}

func dogHandler(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "yo DOG wutup dis is max")	
}

func meHandler(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "yo ME wutup dis is max")	
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/dog/", dogHandler)
	http.HandleFunc("/me/", meHandler)
	http.ListenAndServe(":8080", nil);
}