package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const PROJECTDIRNAME = "inventory-tracking"

func ParseBody(request *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(request.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func loadEnv() {
	projectName := regexp.MustCompile(`^(.*` + PROJECTDIRNAME + `)`)
	currentWorkDir, err := os.Getwd()
	//add error handling
	if err != nil {

	}
	//Root path of the project
	rootPath := string(projectName.Find([]byte(currentWorkDir)))
	err = godotenv.Load(rootPath + "/.env")
}

func GetEnvVariable(key string) string {
	loadEnv()
	return os.Getenv(key)
}
