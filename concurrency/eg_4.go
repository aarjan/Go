// Sample program that checks if the values are equal in the given tree
package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	return false
}

func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Left == nil && t.Right == nil {
		return
	}
	if t.Right != nil {
		Walk(t.Right, ch)
	}
	if t.Left != nil {
		Walk(t.Left, ch)
	}
}

func main() {
	c := make(chan int)
	// x1, x2 := make([]int, 10)
	t1 := tree.New(5)
	t2 := tree.New(5)
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println("--------------")
	go func() {
		for {
			fmt.Println(<-c)
		}
	}()

	Walk(t1, c)

	// Can't use this, because we don't know when the channel is closed
	// for v := range ch {
	// 	fmt.Println(v)
	// }
}
