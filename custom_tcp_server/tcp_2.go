// Sample program to write to the server
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Writes to server through conn pool
	fmt.Fprint(conn, "I dialed you")
}

/*
Inorder to see the request in the server, run the server using tcp_0.go file.
*/
