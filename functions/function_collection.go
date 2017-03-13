package main

import (
	"fmt"
	"math/rand"
	"time"
)

type funcType func(int, int) int

func func1(a funcType) int {
	return a(2, 3)
}

func div(i, j int) int { return i / j }

func main() {
	// Seed your random number generator
	rand.Seed(time.Now().UnixNano())

	a := []funcType{
		func(i, j int) int { return i + j },
		func(i, j int) int { return i * j },
		func(i, j int) int { return i - j },
		funcType(div),	// Type Conversion
	}

	for index := 0; index < len(a); index++ {
		b := func1(a[rand.Intn(len(a))])
		fmt.Println(b)
	}
}
