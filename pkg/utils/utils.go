package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

//PROJECTNAME refers to the project name (base directory).
const PROJECTNAME string = "inventory-tracking"

func ParseBody(request *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(request.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

//Loads the environment variables from .env file in root folder.
func loadEnv() {
	projectName := regexp.MustCompile("^(.*" + PROJECTNAME + ")")
	currentWorkDir, err := os.Getwd()
	//add error handling
	if err != nil {

	}
	//Root path of the project
	rootPath := string(projectName.Find([]byte(currentWorkDir)))
	err = godotenv.Load(rootPath + "/.env")
}

//GetEnvVariable takes a key as parameter and returns the associated value from environment variables.
func GetEnvVariable(key string) string {
	loadEnv()
	return os.Getenv(key)
}
