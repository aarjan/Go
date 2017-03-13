// Sample program running a tcp server at port 8080
// For every connection request, it responds with a message
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintln(conn, "Hello Client! This is your server responding.")
		
        conn.Close()
	}
}
