package models

import (
	"mesh-models-api/Godeps/_workspace/src/gopkg.in/mgo.v2/bson"
	"time"
)

type Item struct {
	Id        bson.ObjectId `bson:"_id,omitempty" json:"id"`
	UID       string        `json:"uid"`
	Url       string        `json:"url"`
	Title     string        `json:"title"`
	CreatedAt time.Time     `json:"created_at"`
}

type Items []Item
