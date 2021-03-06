package models

import (
	"log"
	"mesh-models-api/Godeps/_workspace/src/gopkg.in/mgo.v2/bson"
	"mesh-models-api/connection"
	"mesh-models-api/helpers"
	"time"
)

func Models() []Model {
	var models []Model
	c, session := connection.GetCollection("models")
	defer session.Close()
	err := c.Find(nil).All(&models)
	if err == nil {
		return models
	} else {
		return models
	}
}

func FindModelById(id string) Model {

	c, session := connection.GetCollection("models")
	defer session.Close()
	result := Model{}
	err := c.Find(bson.M{"_id": id}).One(&result)
	if err != nil {
		return result
		log.Fatal(err)
	}

	return result
}

func (o Model) CreateModel() Model {

	c, session := connection.GetCollection("models")
	defer session.Close()
	// This guarantees that t.Id has the bson _id generated by mongo
	o.CreatedAt = time.Now()

	// This while loop ensures there is duplicate Hash
	for {

		id := helpers.RandomString(9)
		count, _ := c.Find(bson.M{"_id": id}).Count()
		if count > 0 {

		} else {

			o.Id = id
			break
		}
	}

	err := c.Insert(o)
	if err != nil {
		return Model{}
		log.Fatal(err)
	}
	return o
}

func DestroyModel(id string) error {

	c, _ := connection.GetCollection("models")

	return c.Remove(bson.M{"_id": id})
}