package main

import (
	"log"
	"net/http"
	"os"
	// "gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	// "fmt"
)

var (

	session *mgo.Session
	err error 
	// db *mgo.Database
)

func getCollection(collectionName string) *mgo.Collection {

	return session.Copy().DB("timlink_development").C(collectionName)
}
// var c *mgo.Collection

func main() {
	session, err = mgo.Dial("localhost:27017")
	// db = session.DB("timlink_development")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))

}
