// Sample program to concatenate template
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	name := "Bond. "
	fullname := "James, Bond"
	text := `<html>
        <body>
            <h1>Hello! I am ` + name + fullname + `</h1>
        </body>
    </html>`

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, text)
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
