package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"timlink/Godeps/_workspace/src/github.com/gorilla/mux"
	"timlink/models"
	// "timlink/connection"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func ItemIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	items := models.RepoAllItems()
	if len(items) != 0 {
		if err := json.NewEncoder(w).Encode(items); err != nil {
			panic(err)
		}
	} else {
		w.Write([]byte("[]"))
	}
}

func ItemShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := string(vars["itemId"])
	item := models.RepoFindItem(itemId)

	if item.Id != "" {

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(item)
	} else {

		// If we didn't find it, 404
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Item Not Found"}); err != nil {
			panic(err)
		}
	}
}

func ItemCreate(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	fmt.Println(item)

	name := r.URL.Query().Get("name")
	url := r.URL.Query().Get("url")

	if name != "" && url != "" {
		item.Name = string(name)
		item.Url = string(url)

		t := models.RepoCreateItem(item)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(t); err != nil {
			panic(err)
		}
	} else {
		code := http.StatusNotFound
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(code)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: code, Text: "url and name param required"}); err != nil {
			panic(err)
		}

	}
}

func ItemDestroy(w http.ResponseWriter, r *http.Request) {
	fmt.Println("adsad")
	vars := mux.Vars(r)
	itemId := string(vars["itemId"])
	err := models.RepoDestroyItem(itemId)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err == nil {

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`"message": "item successfully removed"}`))
		return
	} else {

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Error removing item"}); err != nil {
			panic(err)
		}
		return
	}
	// if item.Id != "" {

	// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// 	w.WriteHeader(http.StatusOK)
	// 	json.NewEncoder(w).Encode(item)
	// } else {

	// 	// If we didn't find it, 404
	// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// 	w.WriteHeader(http.StatusNotFound)
	// 	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Item Not Found"}); err != nil {
	// 		panic(err)
	// 	}
	// }
}
