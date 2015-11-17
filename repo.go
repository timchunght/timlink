package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"log"
	"reflect"
)

var currentId int

var items Items

// Give us some seed data
func init() {

	// c := mongodbSession.Copy().DB("timlink_development").C("items")

	// err = c.Insert(&Item{Name: "Write presentation mongo"},
	// 	&Item{Name: "Host meetup mongo"})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// RepoCreateItem(Item{Name: "Write presentation"})
	// RepoCreateItem(Item{Name: "Host meetup"})
}

func RepoFindItem(id int) Item {
	c := getCollection("items")
	fmt.Println(reflect.TypeOf(c))
	
	result := Item{}
	err = c.Find(bson.M{"name": "helloworld"}).One(&result)
	if err != nil {
		return result
		log.Fatal(err)
	}
	fmt.Println(result.Name)
	return result
}

//this is bad, I don't think it passes race condtions
func RepoCreateItem(t Item) Item {
	c := getCollection("items")

	err = c.Insert(t)
	if err != nil {
		return Item{}
		log.Fatal(err)
	}
	return t
}

// func RepoDestroyItem(id int) error {
// 	for i, t := range items {
// 		if t.Id == id {
// 			items = append(items[:i], items[i+1:]...)
// 			return nil
// 		}
// 	}
// 	return fmt.Errorf("Could not find Item with id of %d to delete", id)
// }
