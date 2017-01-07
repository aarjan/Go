//WAP to set environment varibles
//Environment variables are a universal mechanism for conveying configuration information to Unix programs.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	//To set a key/value pair, use os.Setenv
	os.Setenv("FOO", "1")
	fmt.Println("Foo:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR")) //This will return an empty string if the key isn't present

	fmt.Println()

	for _, e := range os.Environ() { //os.Environ() lists all the key/value paris in the environment
		pair := strings.Split(e, "=") //splits the key and value
		fmt.Println(pair[0])          //prints all the keys
	}
}

// Running the program shows that we pick up the value for FOO that we set in the program, but that BAR is empty.
// $ go run environment-variables.go
// FOO: 1
// BAR:

// The list of keys in the environment will depend on your particular machine.
// TERM_PROGRAM
// PATH
// SHELL
// ...

// If we set BAR in the environment first, the running program picks that value up.
// $ BAR=2 go run environment-variables.go
// FOO: 1
// BAR: 2
