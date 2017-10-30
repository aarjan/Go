package main

import (
	"fmt"
	"sync"
	"time"
)

type SyncCounter struct {
	counter map[string]int
	sync.Mutex
}

func (s *SyncCounter) Inc(key string) {
	s.Lock()
	defer s.Unlock()
	s.counter[key]++
	fmt.Println(s.counter)
}

func (s *SyncCounter) Value(key string) int {
	s.Lock()
	defer s.Unlock()
	return s.counter[key]
}

func main() {
	count := SyncCounter{counter: make(map[string]int)}
	for i := 0; i <= 10; i++ {
		go count.Inc("increase")
	}
	time.Sleep(1000)
	fmt.Println(count.Value("increase"))

}
