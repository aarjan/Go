//Sample program to read in records from a CSV file

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {

	//Open the iris dataset file
	file, err := os.Open("/home/aarzan/gocode/src/github.com/aarjan/golang-training/data/iris.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	//assign a CSV reader for reading from the opened file
	reader := csv.NewReader(file)

	// Assume we don't know the number of fields per line.  By setting
	// FieldsPerRecord negative, each row may have a variable
	// number of fields.
	reader.FieldsPerRecord = -1

	//read all the contents of file
	rawCSVData, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	for _, data := range rawCSVData {
		fmt.Println(data)
	}

}
