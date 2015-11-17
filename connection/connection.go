package connection

import (
	"gopkg.in/mgo.v2"
	"log"

)

var (
	Session *mgo.Session
	// Database *mgo.Database
	err error
)

func Connect() {
	host := "localhost:27017"
	Session, err = mgo.Dial(host)
	if err != nil {
		panic(err)
	}
	Session.SetMode(mgo.Monotonic, true)
	log.Printf("%s\t%s", "connected to", host)
}
