package connection

import (
	"log"
	"mesh-models-api/Godeps/_workspace/src/github.com/joho/godotenv"
	"mesh-models-api/Godeps/_workspace/src/gopkg.in/mgo.v2"
	"os"
)

type DbConfig struct {
	Host         string
	Database     string
	Username     string
	Password     string
	TestDatabase string
	Url          string
}

var (
	Session *mgo.Session
	DbConf  DbConfig
	// Database *mgo.Database
	err error
)

func Connect() {

	DbConf.parseConfig()
	url := DbConf.Url
	Session, err = mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	Session.SetMode(mgo.Monotonic, true)
	log.Printf("%s\t%s", "connected to", url)
}

func GetCollection(collectionName string) (*mgo.Collection, *mgo.Session) {
	session := Session.Copy()
	return session.DB(DbConf.Database).C(collectionName), session
}

func (dbConf *DbConfig) parseConfig() {

	if os.Getenv("GO_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	// Expected Url to be in the following format
	// "mongodb://localhost:27017"
	// mongodb://myuser:mypass@localhost:40001,otherhost:40001/mydb

	// To Be Implemented
	// dbConf := DbConfig{Host: os.Getenv("MONGO_HOST"),
	// 									 Database: os.Getenv("MONGO_DB"),
	// 									 Username: os.Getenv("MONGO_USERNAME"),
	// 									 Password: os.Getenv("MONGO_PASSWORD"),
	// 									 TestDatabase: os.Getenv("MONGO_TEST_DB")}

	dbConf.Url = os.Getenv("MONGO_URL")
	if dbConf.Url == "" {
		dbConf.Url = "mongodb://localhost:27017/mesh-models-api-dev"
	}

}
