package main

import (
	"log"
	"net/http"
	"os"

	"gopkg.in/mgo.v2"
)

var mongodbSession *mgo.Session
var err interface{}

func main() {
	mongodbSession, err = mgo.Dial("localhost:27017")

	if err != nil {
		panic(err)
	}

	mongodbSession.SetMode(mgo.Monotonic, true)

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))

}
