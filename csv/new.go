package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("/home/aarzan/Downloads/Food_Inspections.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	r := csv.NewReader(f)
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rows[0])
	fmt.Println(rows[1])
	

	if err != nil {
		fmt.Println(err)
	}

	// for _,row := range rows{
	// 	fmt.Println(row)
	// }
}
