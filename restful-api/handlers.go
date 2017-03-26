package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello from gorilla ")
}

// TodoIndex ...
func TodoIndex(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(todos)
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

// TodoCreate ...
func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	body, err := ioutil.ReadAll(r.Body)
	checkError(err)
	err = r.Body.Close()
	checkError(err)
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)

		// Send Unmarshal error to user
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatal(err)
		}
	}

	t := RepoCreateTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	// Send the created todo back to user
	if err := json.NewEncoder(w).Encode(t); err != nil {
		log.Fatal(err)
	}
}

// TodoDestroy ...
func TodoDestroy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "please enter an integer", http.StatusBadRequest)
		return
	}

	if err := RepoDestroyTodo(id); err != nil {
		http.Error(w, err.Error(), 422)
		return
	}

	fmt.Fprintf(w, "sucessfully deleted todo of id %d", id)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
