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

// TestBaseHandler holds a gorm database connection.
type TestBaseHandler struct {
	DB *gorm.DB
}

// CreateTestDB returns an in-memory sqlite database connection used for testing.
func CreateTestDBHandler() *TestBaseHandler {
	db := CreateGormDB()
	return &TestBaseHandler{DB: db}
}

func CreateGormDB() *gorm.DB {
	const DBConnection = "file::memory:?cache=shared"
	db, _ := gorm.Open(sqlite.Open(DBConnection), &gorm.Config{})
	db.AutoMigrate(&Item{})
	return db
}
