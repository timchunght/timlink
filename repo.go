package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

var currentId int

var items Items

// Give us some seed data
func init() {
	session, err := mgo.Dial("localhost:27017")

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("timlink_development").C("items")

	err = c.Insert(&Item{Name: "Write presentation mongo"},
								 &Item{Name: "Host meetup mongo"})
	if err != nil {
		log.Fatal(err)
	}


	RepoCreateItem(Item{Name: "Write presentation"})
	RepoCreateItem(Item{Name: "Host meetup"})
}

func RepoFindItem(id int) Item {
	session, err := mgo.Dial("localhost:27017")

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("timlink_development").C("items")

	result := Item{}
	err = c.Find(bson.M{"name": "Write presentation mongos"}).One(&result)
	if err != nil {
		return result
		log.Fatal(err)
	}
	fmt.Println(result.Name)
	return result
}

//this is bad, I don't think it passes race condtions
func RepoCreateItem(t Item) Item {
	currentId += 1
	t.Id = currentId
	items = append(items, t)
	return t
}

func RepoDestroyItem(id int) error {
	for i, t := range items {
		if t.Id == id {
			items = append(items[:i], items[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Item with id of %d to delete", id)
}
