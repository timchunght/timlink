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
	Database *mgo.Database
	err error
)

// Connect to database and set global variable Session and Database
func ConnectDb() {
	Session, err = mgo.Dial("localhost:27017")
	Database = Session.DB("timlink_development")
	if err != nil {
		panic(err)
	}
	Session.SetMode(mgo.Monotonic, true)
}

