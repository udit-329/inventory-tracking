package controllers

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"inventory-tracking/pkg/models"
	"inventory-tracking/pkg/utils"
)

//AddItem adds a new item to the database.
func AddItem(w http.ResponseWriter, r *http.Request) {
	createItem := models.Item{}
	utils.ParseBody(r, &createItem)

	item := createItem.CreateItem()
	res, err := json.Marshal(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Write(res)
}

//GetItemByID fetches an item from the database by its ID.
func GetItemByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, _ := strconv.ParseInt(idStr, 0, 0)

	//Don't get the database.
	details, _ := models.GetItemByID(id)
	res, err := json.Marshal(details)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

//GetItems fetches all the items from the database.
func GetItems(w http.ResponseWriter, r *http.Request) {
	allItems := models.GetItems()
	res, err := json.Marshal(allItems)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

//UpdateItem updates an item in the database.
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	updateItem := models.Item{}
	utils.ParseBody(r, &updateItem)

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, _ := strconv.ParseInt(idStr, 0, 0)

	itemDetails, db := models.GetItemByID(id)

	//Check for updated fields.
	if updateItem.Name != "" {
		itemDetails.Name = updateItem.Name
	}
	if updateItem.Location != "" {
		itemDetails.Location = updateItem.Location
	}
	if updateItem.Quantity != 0 {
		itemDetails.Quantity = updateItem.Quantity
	}

	db.Save(&itemDetails)

	res, err := json.Marshal(itemDetails)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

//DeleteItem deletes an item from the database by its ID.
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, _ := strconv.ParseInt(idStr, 0, 0)

	item := models.DeleteItem(id)

	res, err := json.Marshal(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

//ExportItems generates and exports a CSV file with all product data.
func ExportItems(w http.ResponseWriter, r *http.Request) {
	allItems := models.GetItems()

	//We Marshal and Unmarshal into a generic list of maps to get access to id and creation/updation dates.
	res, _ := json.Marshal(allItems)
	var resUnmarshal []map[string]interface{}
	json.Unmarshal(res, &resUnmarshal)

	buffer := &bytes.Buffer{}
	writer := csv.NewWriter(buffer)

	//Add file headers.
	var headers []string
	headers = append(headers, "Id")
	headers = append(headers, "Name")
	headers = append(headers, "Quantity")
	headers = append(headers, "Location")
	headers = append(headers, "Date Created")
	headers = append(headers, "Date Updated")
	writer.Write(headers)
	for _, item := range resUnmarshal {
		var record []string

		//The ID is originally stored as a float64 type, even though it is basically an it.
		//We can safely cast it to an int64, then convert to a String. Same thing is done with Quantity.
		record = append(record, strconv.FormatInt(int64(item["ID"].(float64)), 10))
		record = append(record, item["Name"].(string))
		record = append(record, strconv.FormatInt(int64(item["Quantity"].(float64)), 10))
		record = append(record, item["Location"].(string))

		//parse and convert time to a more readable format.
		parsedTime, _ := time.Parse(time.RFC3339, item["CreatedAt"].(string))
		record = append(record, parsedTime.Format(time.RFC1123))

		parsedTime, _ = time.Parse(time.RFC3339, item["UpdatedAt"].(string))
		record = append(record, parsedTime.Format(time.RFC1123))

		err := writer.Write(record)
		if err != nil {

		}
	}
	writer.Flush()

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Add("Content-Disposition", "attachment; filename=wow.csv")
	w.Write(buffer.Bytes())
}
