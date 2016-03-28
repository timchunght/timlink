package main

import (

	"net/http"
	"mesh-models-api/controllers"
)

type Route struct {
	Name        string
	Methods     []string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		[]string{"GET"},
		"/",
		controllers.Index,
	},
	Route{
		"ModelCreate",
		[]string{"GET", "POST"},
		"/models",
		controllers.ModelCreate,
	},
	Route{
		"ModelShow",
		[]string{"GET"},
		"/models/{id}",
		controllers.ModelShow,
	},

}
