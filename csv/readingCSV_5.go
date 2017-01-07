// Sample program to read all records of a CSV file,
// and catch an unexpected type in any of the columns.

package main

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
)

type CSVRecord struct {
	SepalLength float64
	SepalWidth  float64
	PetalLenght float64
	PetalWidth  float64
	Species     string
	ParseError  error
}

func main() {

	file, err := os.Open("/home/aarzan/gocode/src/github.com/aarjan/data/iris_mixed_types.csv")

	if err != nil {
		log.Fatal(err)
	}

	// Use this "defer" statement, just how i forgot it, and my whole computer just crashed while running this program
	defer file.Close()

	reader := csv.NewReader(file)

	line := 1
	var csvData []CSVRecord
	var csvRecord CSVRecord

	for {

		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		for idn, value := range record {
			if idn == 4 {
				if value == "" {
					log.Printf("Parsing line %d failed, unexpected type in coulumn %d\n", line, idn)
					csvRecord.ParseError = errors.New("Empty string value")
					break
				}

				// Since any other value is considered as string in csv file
				// Add the string value to the csvRecord
				csvRecord.Species = value
				continue
			}

			// Otherwise, parse the value in record as float64
			var floatValue float64

			if floatValue, err = strconv.ParseFloat(value, 64); err != nil {
				log.Printf("Parsing line %d failed, unexpected type in coulumn %d\n", line, idn)
				csvRecord.ParseError = errors.New("Unexpected value")
				break
			}

			switch idn {
			case 0:
				csvRecord.SepalLength = floatValue
			case 1:
				csvRecord.SepalWidth = floatValue
			case 2:
				csvRecord.PetalLenght = floatValue
			case 3:
				csvRecord.PetalWidth = floatValue
			}

		}

		if csvRecord.ParseError == nil {
			csvData = append(csvData, csvRecord)
		}

		line++

	}

	log.Printf("Succesfully parsed %d rows \n", len(csvData))
}
