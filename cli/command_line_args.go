package main

//os.Args provides access to raw command-line arguments.

import (
	// "bytes"
	"fmt"
	"os"
)

func main() {

	argsWithProg := os.Args
	//Note that the first value in this slice is the path to the program.
	argsWithoutProg := os.Args[1:] //os.Args[1:] holds the arguments to the program.

	arg := os.Args[2]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}

//Build the program and supply the arguments

// go bulid command_line_args.go
// ./command_line_args a b c d
// b
