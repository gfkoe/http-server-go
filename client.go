package main

import (
	"log"
	"net"
)

func main() {
	// Connect to this server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	_, err = conn.Write([]byte("Hello, Server!"))
	if err != nil {
		log.Fatal(err)
	}
}
