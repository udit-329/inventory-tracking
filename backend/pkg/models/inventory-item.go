package models

import (
	"gorm.io/gorm"
)

//Item is a struct that represents an item in the inventory tracking database.
//It stores the item's name, quantity and location where its stored at.
type Item struct {
	gorm.Model
	Name     string
	Quantity int64
	Location string
}

//Functions for Item object.

//CreateItem adds a new Item to the database.
func (item *Item) CreateItem(db *gorm.DB) Item {
	db.Create(&item)
	return *item
}

//GetItemByID takes in an item id as parameter and returns the associated object.
func GetItemByID(db *gorm.DB, id int64) (item Item, dbNew *gorm.DB) {
	dbNew = db.Where("id = ?", id).Find(&item)
	return item, dbNew
}

//GetItems returns a list of all items in the database.
func GetItems(db *gorm.DB) (itemsList []Item) {
	db.Find(&itemsList)
	return itemsList
}

//DeleteItem takes in an item id as parameter and deletes the associated entry from the database.
func DeleteItem(db *gorm.DB, id int64) (item Item) {
	//We also want to return the content of deleted entry.
	db.Where("id = ?", id).Find(&item)
	db.Where("id = ?", id).Delete(&item)
	return item
}
