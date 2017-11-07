package main

import (
	"fmt"
	"sync"
)

/*
	Generator generates the sequence till the range completes.
	The loop is run in a separate goroutine, as we have passed the values to a unbuffered channel.
	If we had run the loop in the same goroutine, there will be deadlock as the channel will get multiple send request.
*/
func generator(start, stop int) <-chan int {
	c := make(chan int)
	go func() {
		for i := start; i < stop; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

/* The fan in pattern is an important pattern which combines
mulitple channels, returns a single channel from those channels
*/

func fanIn(chans ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	// Output channel
	c := make(chan int)

	// Add no. of channels to the waitgroup
	wg.Add(len(chans))

	// For each channel, send the value from the input channel to the output channel
	output := func(ch <-chan int) {
		for n := range ch {
			c <- n
		}
		wg.Done()
	}

	// For each channel, run output in separate go routine
	for _, ch := range chans {
		go output(ch)
	}

	// Wait for all the goroutines before channel is closed
	go func() {
		wg.Wait()
		close(c)
	}()
	return c
}

func main() {
	s1 := generator(1, 10)
	s2 := generator(20, 30)
	s3 := generator(40, 50)
	s4 := generator(60, 70)

	// merge all the channels into one
	mergerd := fanIn(s1, s2, s3, s4)

	for n := range mergerd { // range loop terminates once the chan is closed, otherwise it blocks if there is no value
		fmt.Println(n)
	}
}
