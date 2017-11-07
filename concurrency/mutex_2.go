package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	count = 25
	lock  sync.Mutex
)

func main() {
	for i := 0; i < 20; i++ {
		fmt.Println(i)
		go counter()
		go grade()

	}

	time.Sleep(time.Nanosecond * 1)
	fmt.Println("done")
}

// A mutex serializes access to the code under lock.
// The reason we simply define our lock as lock sync.Mutex is because
// the default value of a sync.Mutex is unlocked.

func counter() {
	lock.Lock()         //locks the goroutine
	defer lock.Unlock() //releases the lock
	count++             //increases the counter in one goroutine at a time
	fmt.Println(count)
}

func grade() {
	lock.Lock()
	defer lock.Unlock()
	count++
	//because of Mutex, the counter is incremented sequentially
	fmt.Println("grade", count)

}
