package main

import (
	"log"
	"net/http"
)

func main() {
	m := NewRouter()
	log.Fatal(http.ListenAndServe(":8000", m))
}
