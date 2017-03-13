// Sample program to implement a rot13 encrypter
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	checkErr(err)
	defer li.Close()

	for {
		conn, err := li.Accept()
		checkErr(err)
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		str := strings.ToLower(scanner.Text())
		buf := []byte(str)
		rot := make([]byte, len(buf))
		for i, v := range buf {
			if v >= 109 {
				rot[i] = v - 13
			} else {
				rot[i] = v + 13
			}
		}
		fmt.Fprint(conn, string(rot))
	}
}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
