// Sample program to illustrate passing function as value

package main

import (
	"fmt"
)

func main() {
	f := func(i int) int {
		return i * i
	}
	m := []int{1, 2, 3, 4}

	fmt.Println(Map(f, m))
}

func Map(f func(int) int, l []int) []int {
	s := make([]int, len(l))

	for n, val := range l {
		s[n] = f(val)
	}
	return s
}
