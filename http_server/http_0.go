package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo ran http.Error(w, err.Error(), http.StatusInternalServerError)")
}

func dog(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("simple_http_server/dog.html")
	if err != nil {
		http.Error(w, "cannot execute template", 404)
	}
	tmpl.ExecuteTemplate(w, "dog.html", nil)

}

func pumpu(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Its a dog")
	http.ServeFile(w, r, "simple_http_server/dog.jpg")
}
func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog.jpg", pumpu)
	http.ListenAndServe(":8080", nil)
}
