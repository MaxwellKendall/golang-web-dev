package main

import (
	"io"
	"log"
	"net"
)

func main(){
	listener, err := net.Listen("tcp", ":8080")
	defer listener.Close()
	if err != nil {
		log.Fatal("There was some error with this program", err)
	}

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		
		io.WriteString(connection, "I see you connected")

		connection.Close()
	}
}