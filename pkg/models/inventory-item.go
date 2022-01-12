package models

import (
	"inventory-tracking/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

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

func (item *Item) CreateItem() *Item {
	db.Create(&item)
	return item
}
