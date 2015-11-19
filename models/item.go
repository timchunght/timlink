package models

import (
	"time"
	"timlink/Godeps/_workspace/src/gopkg.in/mgo.v2/bson"
)

type Item struct {
	Id        bson.ObjectId `bson:"_id,omitempty" json:"id"`
	UID       string        `json:"uid"`
	Url       string        `json:"url"`
	Title   	string        `json:"title"`
	CreatedAt time.Time     `json:"created_at"`
}

type Items []Item
