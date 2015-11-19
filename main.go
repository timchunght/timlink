package main

import (
	"log"
	"net/http"
	"os"
	"timlink/Godeps/_workspace/src/gopkg.in/mgo.v2"
	"timlink/connection"
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



