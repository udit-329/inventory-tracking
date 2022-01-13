package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//ParseBody takes in an http request and a pointer to an object, and parses the request body.
//The parsed JSON body is stored in the original object whose reference is passed.
func ParseBody(r *http.Request, obj interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	json.Unmarshal([]byte(body), obj)
}
