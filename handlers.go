package main

import (
	"encoding/json"
	"fmt"
	"mesh-models-api/Godeps/_workspace/src/github.com/gorilla/mux"
	"mesh-models-api/models"
	"net/http"
	// "mesh-models-api/connection"
)

func Index(w http.ResponseWriter, r *http.Request) {
	welcome_message := `Welcome! This is the simpliest url shortener API

POST   /shorten?url=your_url
Sample Response:
{
  "id": "564d422037f69f505565dc64",
  "url": "https://www.google.com",
  "hash": "dayYGs",
  "short_url": "localhost:9000/dayYGs",
  "created_at": "2015-11-18T22:29:36.900194106-05:00"
}
											`
	fmt.Fprint(w, welcome_message)
}

// func ItemIndex(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.WriteHeader(http.StatusOK)
// 	items := models.RepoAllItems()
// 	if len(items) != 0 {
// 		if err := json.NewEncoder(w).Encode(items); err != nil {
// 			panic(err)
// 		}
// 	} else {
// 		w.Write([]byte("[]"))
// 	}
// }

// func ItemShow(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	itemId := string(vars["itemId"])
// 	item := models.RepoFindItem(itemId)

// 	if item.Id != "" {

// 		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 		w.WriteHeader(http.StatusOK)
// 		json.NewEncoder(w).Encode(item)
// 	} else {

// 		// If we didn't find it, 404
// 		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 		w.WriteHeader(http.StatusNotFound)
// 		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Item Not Found"}); err != nil {
// 			panic(err)
// 		}
// 	}
// }

// func ItemCreate(w http.ResponseWriter, r *http.Request) {
// 	var item models.Item
// 	// fmt.Println(item)

// 	url := r.URL.Query().Get("url")

// 	if url != "" {
// 		item.Url = string(url)

// 		t := item.RepoCreateItem()
// 		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 		w.WriteHeader(http.StatusCreated)
// 		if err := json.NewEncoder(w).Encode(t); err != nil {
// 			panic(err)
// 		}
// 	} else {
// 		code := http.StatusNotFound
// 		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 		w.WriteHeader(code)
// 		if err := json.NewEncoder(w).Encode(jsonErr{Code: code, Text: "url and name param required"}); err != nil {
// 			panic(err)
// 		}

// 	}
// }

// func ItemDestroy(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Println("adsad")
// 	vars := mux.Vars(r)
// 	itemId := string(vars["itemId"])
// 	err := models.RepoDestroyItem(itemId)
// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	if err == nil {

// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(`"message": "item successfully removed"}`))
// 		return
// 	} else {

// 		w.WriteHeader(http.StatusOK)
// 		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Error removing item"}); err != nil {
// 			panic(err)
// 		}
// 		return
// 	}

// }

func LinkCreate(w http.ResponseWriter, r *http.Request) {
	var link models.Link
	// fmt.Println(item)

	url := r.URL.Query().Get("url")

	if url != "" {
		link.Url = string(url)

		l := link.RepoCreateLink()
		l.ShortUrl = r.Host + "/" + l.Hash
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(l); err != nil {
			panic(err)
		}
	} else {
		code := http.StatusNotFound
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(code)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: code, Text: "url param required i.e. ?url=http://google.com"}); err != nil {
			panic(err)
		}

	}
}

func LinkShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash := string(vars["hash"])
	l := models.RepoFindLink(hash)

	if l.Id != "" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		l.ShortUrl = r.Host + "/" + l.Hash
		json.NewEncoder(w).Encode(l)

	} else {

		// If we didn't find it, 404
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Link Not Found"}); err != nil {
			panic(err)
		}
	}
}

func LinkRedirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash := string(vars["hash"])
	link := models.RepoFindLink(hash)

	if link.Id != "" {
		http.Redirect(w, r, link.Url, 301)

	} else {

		// If we didn't find it, 404
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Link Not Found"}); err != nil {
			panic(err)
		}
	}
}
