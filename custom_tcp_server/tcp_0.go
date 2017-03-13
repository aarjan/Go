// Sample program to run a server

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	// here we bind the server to listen at port 8080
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go handle(conn)
	}
}

// Every incomming connection is handled seperately in different goroutines
// We can read from the connection as well as write back to the connection
func handle(conn net.Conn) {

	err := conn.SetDeadline(time.Now().Add(10 * time.Second)) // ends the connection after a deadline
	if err != nil {
		fmt.Println("Connection closed")
	}
	// here we are reading from connection pool
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		fmt.Fprintf(conn, "I Heard you say %s", line) // writing back to the connection
	}
	defer conn.Close()
	fmt.Println("connection ends")
}
