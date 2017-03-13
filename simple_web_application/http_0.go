package main

import (
	// "fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
	// "strings"
)

var (
	// The ParseFiles function takes any number of string arguments that identify our template files, and
	// parses those files into templates that are named after the base file name.
	templates = template.Must(template.ParseFiles("edit.html", "view.html"))

	validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600) // The octal integer literal 0600, indicates the read-write permissions for the current user only
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}

}
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {

	p, e := loadPage(title)
	if e != nil {
		// If no page found; create new
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {

	p, err := loadPage(title)
	if err != nil {
		// If no page found, create new with the given title
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	if string(body) == "" {
		http.Error(w, "error parsing body", http.StatusInternalServerError)
		return
	}
	p := &Page{
		Title: title,
		Body:  []byte(body),
	}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

/*
	The function template.ParseFiles will read the contents of edit.html and return a *template.Template.
	The method t.Execute executes the template, writing the generated HTML to the http.ResponseWriter.
	The .Title and .Body dotted identifiers refer to p.Title and p.Body.
*/
func renderTemplate(w http.ResponseWriter, templ string, p *Page) {
	err := templates.ExecuteTemplate(w, templ+".html", p)
	if err != nil {
		// The http.Error function sends a specified HTTP response code (in this case "Internal Server Error") and error message.
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {

	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.ListenAndServe(":8000", nil)
}
