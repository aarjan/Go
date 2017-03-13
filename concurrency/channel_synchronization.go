/*
We can use channels to synchronize execution across goroutines.
Here’s an example of using a blocking receive to wait for a goroutine to finish.
*/
package main

import (
	"fmt"
	"time"
)

/*
This is the function we’ll run in a goroutine.
The done channel will be used to notify another goroutine that this function’s work is done.
*/
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true // notify that we'e done
}

func main() {
	done := make(chan bool, 1)

	// Start a worker goroutine, giving it the channel to notify on
	go worker(done)

	// Block until we recieve a notification from the worker on the channel.
	<-done
}

/*
If you removed the <- done line from this program,
the program would exit before the worker even started.
*/
