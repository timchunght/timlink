package main

import "fmt"

var currentId int

var items Items

// Give us some seed data
func init() {
	RepoCreateItem(Item{Name: "Write presentation"})
	RepoCreateItem(Item{Name: "Host meetup"})
}

func RepoFindItem(id int) Item {
	for _, t := range items {
		if t.Id == id {
			return t
		}
	}
	// return empty Item if not found
	return Item{}
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
