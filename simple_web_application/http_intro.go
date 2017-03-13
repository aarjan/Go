package main

import (
	"fmt"
	"net/http"
)

func main() {
	// this function tells the http package to handle all request to the web root ("/") with handler.
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

// this function is of type http.HandleFunc
func handler(w http.ResponseWriter, q *http.Request) {

	// http.ResponseWriter value assembles the HTTP server's response
	// By writing to it we send data to the HTTP client
	fmt.Fprintf(w, "Hi there, I love %s!", q.URL.Path[1:])

	/* http.Request is a data structure, that represents the client HTTP request.
	r.URL.Path is the path component of the request URL.
	The trailing [1:] means "create a sub-slice of Path from the 1st character to the end." This drops the leading "/" from the path name.
	*/
}
