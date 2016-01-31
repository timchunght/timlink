package main

import "net/http"

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
		Index,
	},
	// Route{
	// 	"ItemIndex",
	// 	"GET",
	// 	"/items",
	// 	ItemIndex,
	// },
	// Route{
	// 	"ItemCreate",
	// 	"POST",
	// 	"/items",
	// 	ItemCreate,
	// },
	// Route{
	// 	"ItemShow",
	// 	"GET",
	// 	"/items/{itemId}",
	// 	ItemShow,
	// },
	// Route{
	// 	"ItemDestroy",
	// 	"DELETE",
	// 	"/items/{itemId}",
	// 	ItemDestroy,
	// },
	Route{
		"LinkCreate",
		[]string{"GET", "POST"},
		"/shorten",
		LinkCreate,
	},
	Route{
		"LinkShow",
		[]string{"GET"},
		"/links/{hash}",
		LinkShow,
	},
	Route{
		"LinkRedirect",
		[]string{"GET"},
		"/{hash}",
		LinkRedirect,
	},
}
