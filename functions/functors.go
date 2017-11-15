package main

import (
	"fmt"
)

func magical(slice ...int) func(func(int) int) []int {
	return func(f func(int) int) []int {
		result := make([]int, len(slice))
		for i, n := range slice {
			result[i] = f(n)
		}
		return result
	}
}

func main() {
	f := magical(1, 2, 3, 4)
	afunc := func(a int) int {
		return a * a
	}
	fmt.Println(f(afunc))

}
