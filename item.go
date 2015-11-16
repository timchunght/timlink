package main

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Item struct {
	BID       bson.ObjectId `bson:"_id,omitempty"` 
	Id        int       `json:"id"`
	Url       string    `json:"url"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Items []Item
