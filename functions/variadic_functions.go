package main

import (
	"fmt"
)

func varaidicInt(args ...int) {
	//converts args into a slice of type integer
	fmt.Println("integer variables==========")
	for _, i := range args {
		fmt.Println(i)
	}
}

//interface{} can accept any type of parameters
func varaidic(args ...interface{}) {
	fmt.Println("varities of variables========")
	for _, i := range args {
		fmt.Println(i)
	}
}

func main() {
	varaidicInt(23, 34, 3, 3)
	varaidic(2, 3, 3, "adf")
	varaidic(2.3, 4, "asdf", 32)
}
