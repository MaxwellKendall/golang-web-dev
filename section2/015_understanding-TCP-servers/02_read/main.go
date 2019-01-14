package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
	// by default the "token" which scan is looking for is a line of text
		line := scanner.Text()
		fmt.Println(line)
	}
	defer conn.Close()

	// Code never gets here cause line 29 is an infinite loop
	fmt.Println("Code got here.")
}
