package main

import (
	"fmt"
)

type binFunc func(int, int) int

func add(x, y int) int {
	return x + y
}

func (f binFunc) Error() string {
	return "binFunc error"
}

func main() {
	var err error
	err = binFunc(add)
	fmt.Println(err)
}
