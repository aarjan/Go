/*
	Redifining Error Handling for repetative errors
	1. Create a new type
	2. Wrap the initial value
	3. Wrap the behaviour
	4. Wrap the conditional within the behaviour
	
	More at: https://about.sourcegraph.com/go/go-lift/
*/
package main

import (
	"fmt"
)

func main() {
		
	fmt.Println(divide(10,5,1))
	
	d := &Dividinator{10,false}
	d.divide(0)
	d.divide(5)
	fmt.Println(d)
}

type Dividinator struct {
	answer int
	isZero bool
}

func (d *Dividinator) divide(n int) {
	if n==0 {
		d.isZero = true
		return
	}
	d.answer = d.answer/n
}

func (d *Dividinator) String() string {
	if d.isZero{
		return "cannot divide by zero"
	}
	return fmt.Sprintf("The answer is, %d",d.answer)
}
func divide(a,b,c int) (int,error){
	if b==0 || c==0 {
		return 0,fmt.Errorf("cannot divide by zero")
	}
	return (a/b/c),nil
}
