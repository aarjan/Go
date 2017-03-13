// Sample program to imitate a in-memory reader
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
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

var db = make(map[string]string)

func handleConn(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		st := scanner.Text()
		fl := strings.Fields(st)

		switch fl[0] {
		case "GET":
			key := fl[1]

			if db[key] == "" {
				fmt.Fprintln(conn, "no key value pair")
				continue
			}

			fmt.Fprintln(conn, db[key])

		case "SET":
			if len(fl) != 3 {
				fmt.Fprintln(conn, "usage: SET 'KEY' 'VALUE'")
				continue
			}
			key := fl[1]
			val := fl[2]
			db[key] = val

		case "DEL":
			delete(db, fl[1])

		case `exit`:
			fmt.Fprintln(conn, "exiting...")
			os.Exit(1)
		default:
			fmt.Fprintln(conn, "USE 'GET', 'SET' ,'DEL")
		}
	}
	defer conn.Close()
}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
