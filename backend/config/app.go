package config

import (
	"fmt"
	"os"
	"regexp"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

//PROJECTNAME refers to the project name (base directory).
const PROJECTNAME string = "inventory-tracking"

//Connect establishes a connection to the database.
func Connect() {
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	db = database
}

//GetDB retuns a gorm database.
func GetDB() *gorm.DB {
	return db
}

//Loads the environment variables from .env file in root folder.
func init() {
	projectName := regexp.MustCompile("^(.*" + PROJECTNAME + ")")
	currentWorkDir, err := os.Getwd()
	if err != nil {
		log.Error(err)
	}
	//Root path of the project
	rootPath := string(projectName.Find([]byte(currentWorkDir)))
	err = godotenv.Load(rootPath + "/.env")
	if err != nil {
		log.Error(err)
	}
}
