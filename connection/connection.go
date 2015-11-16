package connection
import (
	// "log"
	// "net/http"
	// "os"
	// "gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

var (
	Session *mgo.Session
	// Database *mgo.Database
	err error
)


func Connect() {
	Session, err = mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	Session.SetMode(mgo.Monotonic, true)
}