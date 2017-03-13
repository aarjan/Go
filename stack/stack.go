// Sample program to implement a Stack
package stack

import (
	"errors"
)

type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

func (s *Stack) Pop() (int, error) {
	i := len(s.items)
	if i == 0 {
		return 0, errors.New("index out of range")
	}
	temp := s.items[i-1]
	s.items = s.items[:i-1]
	return temp, nil
}
