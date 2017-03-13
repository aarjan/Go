// Sample program to parse and execute template
package main

import "text/template"
import "os"

var templ *template.Template

func init() {
	// Must is a wrapper around *template.Template; It panics in case of error
	templ = template.Must(template.ParseGlob("*.gohtml"))
}
func main() {
	templ.ExecuteTemplate(os.Stdout, "superman.gohtml", nil)
}
