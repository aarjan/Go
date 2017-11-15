package main

import (
	"fmt"
)

type intfs interface {
	fmt.Stringer
	Map(fn func(int) int) intfs
	// just convinience function to get the underlying slice
	ints() []int
}

type functor struct {
	slice []int
}

func (f functor) Map(fn func(int) int) intfs {
	for i, n := range f.slice {
		f.slice[i] = fn(n)
	}
	return f
}

func (f functor) ints() []int {
	return f.slice
}

func liftIntSlice(s ...int) intfs {
	return functor{slice: s}
}
func (f functor) String() string {
	return fmt.Sprintf("%#v", f.slice)
}

func main() {

	f := func(i int) int {
		return i + i
	}

	fmt.Println(liftIntSlice(1, 2, 3, 4).Map(f))
}
