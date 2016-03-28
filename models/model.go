package models

import (
	"time"
)

type Model struct {
	Id        string 				`bson:"_id,omitempty" json:"id"`
	Url       string        `json:"url"`
	Path 			string        `json:"path"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
}
