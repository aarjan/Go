// Sample program to serve file using io.Copy

package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", doggie)
	http.HandleFunc("/dog.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func doggie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//this will open the url /dog.jpg which in turn will trigger the dogPic handleFunc
	io.WriteString(w, `<img src="/dog.jpg">`)
}

func dogPic(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("simple_http_server/public/pics/dog.jpg")
	if err != nil {
		http.Error(w, "error opening file", 404)
	}
	defer f.Close()
	io.Copy(w, f)
}
