// Sample program to read in records from an example CSV file.
// and catch an unexpected type in a single column.

package main

import (
	"encoding/csv"
	// "fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("/home/aarzan/gocode/src/github.com/aarjan/data/iris_mixed_types.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	reader.FieldsPerRecord = 5

	//secondCoulumn will hold the float values parsed from the second cloumn of the CSV file
	var secondColumn []float64

	//line will help us keep track of line numbers for logging
	line := 1

	// var rawCSVData [][]string

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		// Let's say that we want to gather the second column in the file
		// and validate that it includes only float values (e.g., because
		// we utilize this as a slice of floats later in our application.
		value, err := strconv.ParseFloat(record[1], 64)

		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type:\n", line)
			continue
		}

		secondColumn = append(secondColumn, value)
		line++
	}
	log.Printf("Succesfully parsed %d lines of the second column", len(secondColumn))

	// fmt.Println(secondColumn)
}
