package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(s string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprint(s, i)
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func main() {
	c := fanIn(boring("annie"), boring("marie"))
	for i := 0; i <= 5; i++ {
		fmt.Println(<-c)
	}
}
