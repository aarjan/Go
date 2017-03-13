// Sample program to parse and execute template

package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// ParseFiles reads the system's root directory; build the file to read the current root directory
	temp, err := template.ParseFiles("james.gohtml")
	checkErr(err)
	temp.Execute(os.Stdout, nil)

	// Here, temp is the pointer to template, so we have called its method instead of function
	temp, err = temp.ParseFiles("superman.gohtml", "jason.gohtml")
	checkErr(err)

	// ExecuteTemplate executes the specified template
	temp.ExecuteTemplate(os.Stdout, "superman.gohtml", nil)

	temp.ExecuteTemplate(os.Stdout, "jason.gohtml", nil)

	// It executes the which ever template it finds at first.
	temp.Execute(os.Stdout, nil)

	temp, err = template.ParseGlob("*.gohtml")
	checkErr(err)
	temp.ExecuteTemplate(os.Stdout, "superman.gohtml", 23)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
