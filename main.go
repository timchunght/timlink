package main

import (
	"log"
	"mesh-models-api/Godeps/_workspace/src/gopkg.in/mgo.v2"
	"mesh-models-api/connection"
	"net/http"
	"os"
)

var (
	session *mgo.Session
	err     error
	// db *mgo.Database
	// col *mgo.Collection
)

func main() {

	connection.Connect()
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))

}
