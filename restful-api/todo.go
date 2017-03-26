package main

import "time"

// Todo ...
type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
	ID        int       `json:"id"`
}

// Todos ... Todo List
type Todos []Todo
