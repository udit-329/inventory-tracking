package config

import (
	"fmt"
	"inventory-tracking/pkg/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	database *gorm.DB
)

func Connect() {
	dbHost := utils.GetEnvVariable("DB_HOST")
	dbName := utils.GetEnvVariable("DB_NAME")
	dbUser := utils.GetEnvVariable("DB_USER")
	dbPass := utils.GetEnvVariable("DB_PASS")
	dbPort := utils.GetEnvVariable("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	database = db
}

func GetDB() *gorm.DB {
	return database
}
