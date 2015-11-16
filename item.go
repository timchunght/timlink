package main

import "time"

type Item struct {
	Id        int       `json:"id"`
	Url       string    `json:"url"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Items []Item
