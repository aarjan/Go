package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func Decode() {
	b := []byte(`{"Name":"Arjan","Age":22,"Address":"Shantinagar","Parents":["ShreeRam","Kalpana"]}`)

	var i interface{}

	err := json.Unmarshal(b, &i)

	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(i)

	expected := map[string]interface{}{
		"Name":    "Arjan",
		"Age":     int(22),
		"Address": "Shantinagar",
		"Parents": []interface{}{
			"ShreeRam",
			"Kalpana",
		},
	}

	for l, n := range expected {
		switch vv := n.(type) {
		case string:
			fmt.Println(l, n)
		case int:
			fmt.Println(l, n)
		case []interface{}:
			fmt.Println(n, "is an array")
			for x, y := range vv {
				fmt.Println(x, y)
			}

		}
	}

}

func main() {
	Decode()
	fmt.Println("Success")
}
