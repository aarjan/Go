package main

import (
	"fmt"
	"github.com/aarjan/counters"
	"github.com/aarjan/users"
)

func main() {
	//returns value of exported type
	counter := counters.AlertCounter(23)
	fmt.Println(counter)

	//returns value of unexported type using exported function
	length := counters.ReturnLength(32.423)
	fmt.Println(length)

	//returns exported fields of embedded struct
	//here Book is the field name in users, users.Book is the exported struct name of the package users
	user1 := users.User{Name: "aarjan", Age: 23, Address: "balkot", Book: users.Book{Title: "Golang nuts", Cost: 23.32}}

	//set the exported fields from the unexported innertype
	user1.PublicKey = 1234
	user1.PrivateKey = "aarzan@123"

	//user1.Book.Title = "adsf"
	fmt.Println(user1)

}
