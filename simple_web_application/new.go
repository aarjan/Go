package main

import (
	"fmt"
	"math/rand"
)

func main() {

	a := rand.Perm(100)

	b := rand.Perm(100)

	c := make([]int, 100)

	// var n int

	for i, j := range a {
		c[i] = j
	}

	for i, j := range b {
		c[i] = c[i] + j
	}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
