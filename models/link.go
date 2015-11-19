package models

import (
	"time"
	"timlink/Godeps/_workspace/src/gopkg.in/mgo.v2/bson"
)

type Link struct {
	Id        bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Url       string        `json:"url"`
	ShortUrl  string        `bson:"short_url" json:"short_url"`
	CreatedAt time.Time     `json:"created_at"`
}

type Links []Link
