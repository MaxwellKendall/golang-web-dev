package main

import (
	"fmt"
	"net/http"
)

type hotdog int

// implicitly implements the Handler Interface
func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Any code you want in this func")
}

func main() {

}
