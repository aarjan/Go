package main

import (
	"fmt"
	//	"io"
	"log"
	"os"
)

func main() {
	open()
}

func open() {
	file, err := os.Open("data/readme.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 137)
	count, err2 := file.Read(data)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Printf("read %d bytes:\n %q ", count, data[:count])
}
