package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello from gorilla ")
}

// TodoIndex ...
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Gate"},
		Todo{Name: "Paint"},
	}
	err := json.NewEncoder(w).Encode(&todos)
	if err != nil {
		http.Error(w, " error in encoding", http.StatusInternalServerError)
	}

}

// TodoShow ...
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintln(w, id)
}
