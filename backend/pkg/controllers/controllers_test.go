package controllers

import (
	"bytes"
	"encoding/json"
	"inventory-tracking/backend/pkg/models"
	"inventory-tracking/backend/pkg/utils"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var TestBaseHandler *BaseHandler

var testItem = models.Item{
	Name:     "Test object",
	Quantity: 5,
	Location: "Test Location",
}

var testItemUpdate = models.Item{
	Name:     "Test object updated",
	Quantity: 10,
	Location: "Test Location updated",
}

func TestMain(m *testing.M) {
	TestBaseHandler = NewBaseHandler(utils.CreateGormDB())

	//Run tests.
	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestAddItem(t *testing.T) {
	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(testItem)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/add", &body)

	TestBaseHandler.AddItem(w, req)

	res := w.Result()
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Error("AddItem failed, got an error.")
	}

	var obj interface{}

	json.Unmarshal([]byte(data), &obj)

	if obj.(map[string]interface{})["Name"] != testItem.Name || int64(obj.(map[string]interface{})["Quantity"].(float64)) != testItem.Quantity || obj.(map[string]interface{})["Location"] != testItem.Location {
		t.Error("AddItem failed, unexpected response.")
	} else if res.StatusCode != 200 {
		t.Error("AddItem failed, did not return 200 status code.")
	}
}

func TestGetItemByID(t *testing.T) {

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/get?id=1", nil)

	TestBaseHandler.GetItem(w, req)

	res := w.Result()
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Error("GetItemByID failed, got an error.")
	}

	var obj interface{}

	json.Unmarshal([]byte(data), &obj)

	if obj.(map[string]interface{})["Name"] != testItem.Name || int64(obj.(map[string]interface{})["Quantity"].(float64)) != testItem.Quantity || obj.(map[string]interface{})["Location"] != testItem.Location {
		t.Errorf("GetItemByID failed, unexpected response.%s", obj)
	} else if res.StatusCode != 200 {
		t.Error("GetItemByID failed, did not return 200 status code.")
	}
}

func TestGetItems(t *testing.T) {

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/get", nil)

	TestBaseHandler.GetItem(w, req)

	res := w.Result()
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Error("GetItems failed, got an error.")
	}

	var obj interface{}

	json.Unmarshal([]byte(data), &obj)
	objItems := obj.([]interface{})[0]

	if objItems.(map[string]interface{})["Name"] != testItem.Name || int64(objItems.(map[string]interface{})["Quantity"].(float64)) != testItem.Quantity || objItems.(map[string]interface{})["Location"] != testItem.Location {
		t.Error("GetItems failed, unexpected response.")
	} else if res.StatusCode != 200 {
		t.Error("GetItems failed, did not return 200 status code.")
	} else if len(obj.([]interface{})) != 1 {
		t.Error("GetItems failed, unexpected number of items returned.")
	}
}

func TestUpdateItem(t *testing.T) {
	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(testItemUpdate)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/update?id=1", &body)

	TestBaseHandler.UpdateItem(w, req)

	res := w.Result()
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Error("UpdateItem failed, got an error.")
	}

	var obj interface{}

	json.Unmarshal([]byte(data), &obj)

	if obj.(map[string]interface{})["Name"] != testItemUpdate.Name || int64(obj.(map[string]interface{})["Quantity"].(float64)) != testItemUpdate.Quantity || obj.(map[string]interface{})["Location"] != testItemUpdate.Location {
		t.Error("UpdateItem failed, unexpected response.")
	} else if res.StatusCode != 200 {
		t.Error("UpdateItem failed, did not return 200 status code.")
	}

	//Fetch the item from the database to check its value.
	req = httptest.NewRequest(http.MethodGet, "/get?id=1", nil)

	TestBaseHandler.UpdateItem(w, req)

	res = w.Result()
	data, err = ioutil.ReadAll(res.Body)

	if err != nil {
		t.Error("UpdateItem failed, got an error.")
	}

	json.Unmarshal([]byte(data), &obj)

	if obj.(map[string]interface{})["Name"] != testItemUpdate.Name || int64(obj.(map[string]interface{})["Quantity"].(float64)) != testItemUpdate.Quantity || obj.(map[string]interface{})["Location"] != testItemUpdate.Location {
		t.Error("UpdateItem failed, unexpected response.")
	} else if res.StatusCode != 200 {
		t.Error("UpdateItem failed, did not return 200 status code.")
	}
}

func TestDeleteItem(t *testing.T) {

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/delete?id=1", nil)

	TestBaseHandler.DeleteItem(w, req)

	res := w.Result()
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Error("DeleteItem failed, got an error.")
	}
	if res.StatusCode != 200 {
		t.Error("DeleteItem failed, did not return 200 status code.")
	}

	//Check the number of items in the database.
	req = httptest.NewRequest(http.MethodGet, "/get", nil)

	TestBaseHandler.GetItem(w, req)

	res = w.Result()
	data, err = ioutil.ReadAll(res.Body)

	var obj interface{}

	if obj != nil {
		t.Errorf("DeleteItem failed, unexpected number of items returned.%s", data)
	}
}
