package main

import (
	"fmt"
	"net"
)

// Define the connection values
const (
	host     = "localhost"
	port     = "1234"
	conntype = "tcp"
)

func main() {
	listener, err := net.Listen(conntype, host+":"+port) // make our listener

	if err != nil {
		fmt.Println(err)
	}

	defer listener.Close()
	fmt.Println("We are now listening on " + host + " and the port of " + port) // little debug message

	for {
		conn, err := listener.Accept() // accept incoming connections
		if err != nil {
			fmt.Println(err) // catch the error
		}
		go handleRequest(conn) //run our handler
	}

}

// handles the incoming request
func handleRequest(conn net.Conn) {
	// buffer holds incoming data
	buf := make([]byte, 1024)
	// reads the data sent
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error()) // catch error
	}
	// tell them there message has been recieved
	conn.Write([]byte("Message received."))
	// close the connection
	conn.Close()
}
