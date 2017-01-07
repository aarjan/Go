package main

import (
	"encoding/json"
	"log"
	"reflect"
)

type Family struct {
	Name   string
	Number int
	Caste  string
}

func Decode() {
	b := []byte(`{"Name":"Baskota","Number":5,"Caste":"Brahmin"}`)
	var m Family
	err := json.Unmarshal(b, &m)

	if err != nil {
		panic(err)
	}

	expected := Family{"Baskota", 5, "Brahmin"}
	if !reflect.DeepEqual(expected, m) {
		log.Panicln("error unmarshalling")
	}

}
func main() {
	Decode()
	print("Success")
}
