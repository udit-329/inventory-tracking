package utils

import (
	"gorm.io/driver/sqlite"
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

type TestBaseHandler struct {
	DB *gorm.DB
}

func CreateTestDB() *TestBaseHandler {
	const DBConnection = "file::memory:?cache=shared"

	db, _ := gorm.Open(sqlite.Open(DBConnection), &gorm.Config{})
	db.AutoMigrate(&Item{})

	return &TestBaseHandler{DB: db}
}
