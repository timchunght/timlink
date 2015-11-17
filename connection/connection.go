package connection

import (

	"timlink/Godeps/_workspace/src/gopkg.in/mgo.v2"
	"log"
	"timlink/Godeps/_workspace/src/gopkg.in/yaml.v2"
	// "fmt"
	"io/ioutil"
  "path/filepath"
)

type DbConfig struct {
  

  Host string   `yaml:"host"`
  Database   string `yaml:"database"`


}

var (
	Session *mgo.Session
	DbConf DbConfig
	// Database *mgo.Database
	err error
)

func Connect() {

	filePath, _ := filepath.Abs("./connection/database.yml")
  databaseFile, err := ioutil.ReadFile(filePath)
	_ = yaml.Unmarshal(databaseFile, &DbConf)
  
	host := DbConf.Host
	
	Session, err = mgo.Dial(host)
	if err != nil {
		panic(err)
	}
	Session.SetMode(mgo.Monotonic, true)
	log.Printf("%s\t%s", "connected to", host)
}

func GetCollection(collectionName string) *mgo.Collection {

	return Session.Copy().DB(DbConf.Database).C(collectionName)
}

