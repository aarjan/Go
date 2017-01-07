// WAP to show flag package supporting basic command-line flag parsing
package main

import (
	"flag"
	"fmt"
)

func main() {
	// Basic flag declarations are available for string,integer and boolean options.
	// flag.String(flag_name, default_value, usage)

	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("number", 35, "an int")

	boolPtr := flag.Bool("fork", false, "is an bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "is an string")

	flag.Parse() // once all flags are declared, flag.Parse() executes the command line Parsing

	//Here we'll just dump out the parsed options and any trailing positional arguments.
	//Note that we need to dereference the pointers eg.*wordPtr to get the actual option values.
	fmt.Println("word:", *wordPtr)
	fmt.Println("number:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())

}

//Note if you omit flags they automatically take their default values

//Trailing positional arguments can be provided after any flags.

// Note that the flag package requires all flag to appear before the positional arguments(otherwise the flags will be interpreted as positional arguments)

// Use -h or --help to get automatically generated usage text

// If you provide a flag that wasn't specified to the flag package, the program will print an error message and show the help text again.

// go build command_line_flags.go
// ./command_line_flags -word=opt -number=32 -fork -svar=fasd
// word : opt
// number : 32
// fork : false
// svar : flag

//Note if you omit flags they automatically take their default values
// $ ./command-line-flags -word=opt
// word: opt
// numb: 42
// fork: false
// svar: bar
// tail: []

// Trailing positional arguments can be provided after any flags.
// $ ./command-line-flags -word=opt a1 a2 a3
// word: opt
// ...
// tail: [a1 a2 a3]

// Note that the flag package requires all flags to appear before positional arguments (otherwise the flags will be interpreted as positional arguments).
// $ ./command-line-flags -word=opt a1 a2 a3 -numb=7
// word: opt
// numb: 42
// fork: false
// svar: bar
// tail: [a1 a2 a3 -numb=7]

// Use -h or --help flags to get automatically generated help text for the command-line program.
// $ ./command-line-flags -h
// Usage of ./command-line-flags:
//   -fork=false: a bool
//   -numb=42: an int
//   -svar="bar": a string var
//   -word="foo": a string

// If you provide a flag that wasn’t specified to the flag package, the program will print an error message and show the help text again.
// $ ./command-line-flags -wat
// flag provided but not defined: -wat
// Usage of ./command-line-flags:
// ...
