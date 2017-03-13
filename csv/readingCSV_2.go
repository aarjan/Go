//Sample program to read the records of a CSV file
//and catch the unexpected extra field

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	// "strings"
)

func main() {
	f, err := os.Open("/home/aarzan/gocode/src/github.com/aarjan/golang-training/data/Employee.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//create a new csv reader from the opened file
	reader := csv.NewReader(f)

	//since we know the expected fields in the record
	reader.FieldsPerRecord = 5

	reader.TrailingComma
	//rawCSVData would our successfully parsed rows
	var rawCSVData [][]string

	//indefinite for loop until a 'break' is encountered
	for {

		// Read in a row
		record, err := reader.Read()

		// s := strings.Split(string(record), "CR")
		// fmt.Println(s)
		// Check if we are in the end of the file
		if err == io.EOF {
			break //breaks from the for loop
		}

		// If we had parsing error, log the error and move on.
		if err != nil {
			log.Println(err)
			continue
		}

		// Append the record to our data, if it has the expected no. of fields
		rawCSVData = append(rawCSVData, record)

	}

	// Outputs the no. of records successfully read
	log.Printf("Successfully parsed %d lines \n", len(rawCSVData))

	// Outputs the data; as a check
	for _, data := range rawCSVData {
		fmt.Println(data)
	}
}
