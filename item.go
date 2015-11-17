package main

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Item struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"id"` 
	UID       string 		`json:"uid"` 
	Url       string    `json:"url"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Items []Item
