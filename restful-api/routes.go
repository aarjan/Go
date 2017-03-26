package main

import "net/http"

// Route ...
type Route struct {
	Name        string
	Method      string
	Pattern     string
	handlerFunc http.HandlerFunc
}

// Routes ...
type Routes []Route

var routes = Routes{

	Route{
		"index",
		"Get",
		"/",
		Index,
	},
	Route{
		"todoIndex",
		"Get",
		"/todos",
		TodoIndex,
	},
	Route{
		"pattern",
		"Get",
		"/todos/{id:[0-9]+}",
		TodoShow,
	},
	Route{
		"todoCreate",
		"Post",
		"/todos/create",
		TodoCreate,
	},
	Route{
		"todoDestroy",
		"Get",
		"/todos/destroy/{id}",
		TodoDestroy,
	},
}
