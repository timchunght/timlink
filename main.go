package main

import (
	"log"
	"net/http"
	"os"
	// "gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	// "fmt"
)

var session *mgo.Session
var err error 
// var c *mgo.Collection

func main() {
	session, err = mgo.Dial("localhost:27017")

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))

}
