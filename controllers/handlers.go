package controllers

import (
	"encoding/json"
	"fmt"
	"mesh-models-api/Godeps/_workspace/src/github.com/gorilla/mux"
	"mesh-models-api/models"
	"mesh-models-api/helpers"
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


func ModelCreate(w http.ResponseWriter, r *http.Request) {
	var model models.Model
	// fmt.Println(item)

	url := r.URL.Query().Get("url")

	if url != "" {
		model.Url = string(url)

		o := model.CreateModel()
		
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(o); err != nil {
			panic(err)
		}
	} else {
		code := http.StatusNotFound
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(code)
		if err := json.NewEncoder(w).Encode(helpers.JsonErr{Code: code, Text: "url param required i.e. ?url=http://google.com"}); err != nil {
			panic(err)
		}

	}
}

func ModelShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := string(vars["id"])
	o := models.FindModelById(id)

	if o.Id != "" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(o)

	} else {

		// If we didn't find it, 404
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(helpers.JsonErr{Code: http.StatusNotFound, Text: "Model Not Found"}); err != nil {
			panic(err)
		}
	}
}

