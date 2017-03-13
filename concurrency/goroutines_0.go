// Sample program to illustrate uses of Goroutines
package main

import (
	"fmt"
)

func f(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
	}
}

func main() {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	/*
		Our two function calls are running asynchronously in separate goroutines now,
		so execution falls through to here.
		Scanln() command  delays the program and lets the goroutines to complete the execution, otherwise the
		the entire program would normally exit without execution of goroutines.
	*/
	var input string
	fmt.Scanln(&input)
	fmt.Println(input, "done")
	
}

/*
When we run this program, we see the output of the blocking call first,
then the interleaved output of the two gouroutines.
This interleaving reflects the goroutines being run concurrently by the Go runtime.
*/
