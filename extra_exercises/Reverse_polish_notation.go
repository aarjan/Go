// Sample program to illustrate Reverse-polish-notation using stack package.

package main

import (
	"fmt"
	"github.com/aarjan/stack"
)

func main() {
	var s stack.Stack

	var input [9]interface{}
	input[0] = 5
	input[1] = 1
	input[2] = 2
	input[3] = '+'
	input[4] = 4
	input[5] = '*'
	input[6] = '+'
	input[7] = 3
	input[8] = '-'

	for _, val := range input {
		switch val.(type) {

		case int:
			t, _ := val.(int)
			s.Push(t)

		case int32:
			var v1, v2 int

			v1, _ = s.Pop()
			v2, _ = s.Pop()

			switch val {
			case '+':
				temp := v2 + v1
				fmt.Println(temp)
				s.Push(temp)
			case '-':
				temp := v2 - v1
				s.Push(temp)
				fmt.Println(temp)
			case '*':
				temp := v2 * v1
				s.Push(temp)
				fmt.Println(temp)
			case '/':
				temp := v2 / v1
				s.Push(temp)
				fmt.Println(temp)
			}
		}
	}
	fmt.Println(s)
}
