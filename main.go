package main

import (
	"log"
	"net/http"
	"os"
	"timlink/connection"
	"gopkg.in/mgo.v2"
)

var (
	session *mgo.Session
	err error 
	// db *mgo.Database
	// col *mgo.Collection
)

func getCollection(collectionName string) *mgo.Collection {

	return connection.Session.Copy().DB("timlink_development").C(collectionName)
}


func main() {
	connection.Connect()
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))

}
