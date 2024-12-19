package main

import (
	"fmt"
	"net"
)

type TCPServer struct {
	Host string
	Port int
}

func (server *TCPServer) RunServer() {
	addr := fmt.Sprintf("%s:%d", server.Host, server.Port)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer listener.Close()

	fmt.Println("Server listening on port:", server.Port)

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	defer connection.Close()
	buffer := make([]byte, 1024)
	_, err := connection.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buffer))
}

func main() {
	server := TCPServer{
		Host: "localhost",
		Port: 8080,
	}
	server.RunServer()
}
