package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("simple_http_server/public/templates/index.html"))
}

func main() {
	http.HandleFunc("/", dogs)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("simple_http_server/public"))))
	http.ListenAndServe(":8080", nil)
}

func dogs(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}
