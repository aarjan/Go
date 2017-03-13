// Sample program that dials to tcp server at port 8080
// Here, the client gets the response from the server and prints it
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	// client dials to the server for connection
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
    defer conn.Close()
    
	// reads the server response
	bs, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))
}
