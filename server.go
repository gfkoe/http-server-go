package main

import (
	"fmt"
	"net"
	"strings"
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

	request := strings.Split(string(buffer), "\n")

	requestParts := strings.Split(request[0], " ")

	header := request[1]

	httpMethod := requestParts[0]

	response := handleRequest(httpMethod, header)

	connection.Write([]byte(response))

	fmt.Println(response)
}

func handleRequest(httpMethod string, header string) string {
	response := ""
	switch httpMethod {
	case "GET":
		response = handleGetRequest(httpMethod, header)
	case "POST":
		response = handlePostRequest(httpMethod, header)
	case "PUT":
		response = handlePutRequest(httpMethod, header)
	case "DELETE":
		response = handleDeleteRequest(httpMethod, header)
	default:
		response = handleErrorRequest(httpMethod, header)
	}
	return response
}

func handleGetRequest(request string, header string) string {
	return "Request: " + request + " " + header
}

func handlePostRequest(request string, header string) string {
	return "Request: " + request + " " + header
}

func handlePutRequest(request string, header string) string {
	return "Request: " + request + " " + header
}

func handleDeleteRequest(request string, header string) string {
	return "Request: " + request + " " + header
}

func handleErrorRequest(request string, header string) string {
	return "404 Error: Invalid HTTP Method: " + request + "\nWith header: " + header
}

func main() {
	server := TCPServer{
		Host: "localhost",
		Port: 8080,
	}
	server.RunServer()
}
