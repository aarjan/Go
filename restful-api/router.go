package main

import "github.com/gorilla/mux"
import "net/http"

// NewRouter ...
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {

		var handler http.Handler
		handler = route.handlerFunc
		handler = Logger(handler, route.Name)

		router.Methods(route.Method).
			Name(route.Name).
			Path(route.Pattern).
			Handler(handler)
	}

	return router
}
