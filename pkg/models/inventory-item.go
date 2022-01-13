package models

import (
	"inventory-tracking/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

//Item is a struct that represents an item in the inventory tracking database.
//It stores the item's name, quantity and location where its stored at.
type Item struct {
	gorm.Model
	Name     string
	Quantity int64
	Location string
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Item{})
}

//Functions for Item object.

//CreateItem adds a new Item to the database.
func (item *Item) CreateItem() Item {
	db.Create(&item)
	return *item
}

//GetItemByID takes in an item id as parameter and returns the associated object.
func GetItemByID(id int64) (item Item, dbNew *gorm.DB) {
	dbNew = db.Where("id = ?", id).Find(&item)
	return item, dbNew
}

//GetItems returns a list of all items in the database.
func GetItems() (itemsList []Item) {
	db.Find(&itemsList)
	return itemsList
}

//DeleteItem takes in an item id as parameter and deletes the associated entry from the database.
func DeleteItem(id int64) (item Item) {
	//We also want to return the content of deleted entry.
	db.Where("id = ?", id).Find(&item)
	db.Where("id = ?", id).Delete(&item)
	return item
}
