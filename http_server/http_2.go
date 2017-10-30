package main

import (
	"html/template"
	"log"
	"net/http"
)

var templ *template.Template

func init() {
	templ = template.Must(templ.ParseGlob("simple_http_server/public/templates/*"))

}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/apply", apply)
	http.HandleFunc("/contact", contact)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	err := templ.ExecuteTemplate(w, "index.html", nil)
	handleError(w, err)
}

func about(w http.ResponseWriter, r *http.Request) {
	err := templ.ExecuteTemplate(w, "about.html", nil)
	handleError(w, err)
}

func apply(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := templ.ExecuteTemplate(w, "apply.html", nil)
		handleError(w, err)
		return
	}
	err := templ.ExecuteTemplate(w, "applyProcess.html", nil)
	handleError(w, err)
}

func contact(w http.ResponseWriter, r *http.Request) {
	err := templ.ExecuteTemplate(w, "contact.html", nil)
	handleError(w, err)
}

func handleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
