package models

import (
	"inventory-tracking/inventory-tracking/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Item struct {
	gorm.model
	Name     string
	Quantity int64
	Location string
}

func init() {
	config.Connect()
	db = config.GetDB()
}
