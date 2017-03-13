/*
Channels are the pipes that connect concurrent goroutines.
They can be used for communicating in/out of goroutines.
You can send values into channels from one goroutine and receive those values into another goroutine.
*/

package main

import (
	"fmt"
)

func main() {
	messages := make(chan string) // channels are typed by the values they contain
	go func() {
		messages <- "ping" // send in to channel
	}()

	msg := <-messages // recieve from channel
	fmt.Println(msg)
}

/*
When we run the program the "ping" message is successfully passed from one goroutine
to another via our channel.

By default senders and receivers block program execution until both the sender and receiver are ready.
This property allowed us to wait at the end of our program for the "ping" message
without having to use any other synchronization.
*/
