package main

import (
	"encoding/json"
	"fmt"
	// "io"
	// "io/ioutil"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func ItemIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(items); err != nil {
		panic(err)
	}
}

func ItemShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var itemId int
	var err error
	if itemId, err = strconv.Atoi(vars["itemId"]); err != nil {
		panic(err)
	}
	item := RepoFindItem(itemId)
	if item.Id != "" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(item)
		// if err := json.NewEncoder(w).Encode(item); err != nil {
		// 	panic(err)
		// }
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

/*
Test with this curl command:

curl -H "Content-Type: application/json" -d '{"name":"New Item"}' http://localhost:8080/items

*/
func ItemCreate(w http.ResponseWriter, r *http.Request) {
	var item Item
	fmt.Println(item)
	
	name := r.URL.Query().Get("name")
	url := r.URL.Query().Get("url")
	// fmt.Println(string(name))
	item.Name = string(name)
	item.Url = string(url)

	t := RepoCreateItem(item)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
