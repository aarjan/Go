package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 2, 3}
	for _, v := range s[2:] {
		fmt.Println(v)
	}
}
