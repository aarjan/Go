package main

import (
	"html/template"
	"log"
	"os"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	slice := []int{2, 3, 35, 4}
	err := tmpl.ExecuteTemplate(os.Stdout, "james.html", slice)
	if err != nil {
		log.Fatal(err)
	}
}
