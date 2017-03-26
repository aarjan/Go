package main

import (
	"fmt"
)

var todos Todos
var currentID int

func init() {

}

// RepoCreateTodo ...
func RepoCreateTodo(t Todo) Todo {
	currentID++
	t.ID = currentID
	todos = append(todos, t)
	return t
}

// RepoFindTodo ...
func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.ID == id {
			return t
		}
	}
	// else return empty Todo
	return Todo{}
}

// RepoDestroyTodo ...
func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("No existing repository for todo with id %d", id)

}
