// Sample program to save records to a CSV file.

package main

import (
	"encoding/csv"
	"log"
	"os"
)

var data = [][]string{
	{"1.2", "1.3", "0.3", "0.12", "Iris-setosa"},
	{"1.0", "2.1", "0.4", "0.8", "Iris-setosa"},
	{"2.1", "8.2", "0.7", "0.2", "Iris-setosa"},
	{"3.2", "1.8", "0.2", "0.15", "Iris-versicolor"},
	{"2.5", "2.7", "0.5", "0.1", "Iris-versicolor"},
	{"1.7", "3.5", "1.0", "0.7", "Iris-virginica"},
	{"1.7", "3.1", "0.5", "0.2", "Iris-virginica"},
	{"1.1", "3.0", "0.2", "0.1", "Iris-virginica"},
}

func main() {

	// Create the output file
	file, err := os.Create("/home/aarzan/gocode/src/github.com/aarjan/data/output.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)

	err = writer.WriteAll(data)
	if err != nil {
		log.Fatal(err)
	}

}
