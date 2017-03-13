/*
By default channels are unbuffered,
meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value.
Buffered channels accept a limited number of values without a corresponding receiver for those values.
Channel sends to the buffered block only when the buffer is full; recieves block when the buffer is empty.
*/

package main

import (
	"fmt"
)

func main() {

	messages := make(chan string, 2) // channel of string takes of to 2 values

	// Because the channel is buffered,
	// we can send these values to the channel without a corresponding concurrent recieve.
	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
