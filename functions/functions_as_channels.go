package main

import (
	"fmt"
	"math/rand"
	"time"
)

type chanFn func(int, int) int

func pickFn(fns ...chanFn) chanFn {
	return fns[rand.Intn(len(fns))]
}

func produce(c chan chanFn, n int, fns ...chanFn) {
	defer close(c)
	for i := 0; i < n; i++ {
		c <- pickFn(fns...)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ch := make(chan chanFn)
	fns := []chanFn{
		func(i, j int) int { return i + j },
		func(i, j int) int { return i * j },
		func(i, j int) int { return i - j },
		func(i, j int) int { return i / j },
	}

	go produce(ch, len(fns), fns...)

	for fn := range ch {

		res := fn(2, 3)
		fmt.Println(res)
		time.Sleep(200 * time.Millisecond) // Delay
	}

}
