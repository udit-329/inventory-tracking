package controllers

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"inventory-tracking/backend/pkg/models"
	"inventory-tracking/backend/pkg/utils"
)

// BaseHandler will hold everything that controller needs
type BaseHandler struct {
	db *gorm.DB
}

//NewBaseHandler returns a new BaseHandler object.
func NewBaseHandler(db *gorm.DB) *BaseHandler {
	db.AutoMigrate(&models.Item{})
	return &BaseHandler{
		db: db,
	}
}

//AddItem adds a new item to the database.
func (handler *BaseHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	createItem := models.Item{}
	utils.ParseBody(r, &createItem)

	//Allow default location.
	if createItem.Location == "" {
		createItem.Location = "N/A"
	}
	//Do not allow empty values.
	if createItem.Name == "" {
		w.WriteHeader(http.StatusInternalServerError)
		res, _ := json.Marshal(map[string]string{"Error": "Name cannot be empty."})
		w.Write(res)
	}
	if createItem.Quantity < 0 {
		w.WriteHeader(http.StatusInternalServerError)
		res, _ := json.Marshal(map[string]string{"Error": "Quantity cannot be less than 0."})
		w.Write(res)
	}

	item := createItem.CreateItem(handler.db)
	res, err := json.Marshal(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Write(res)
}

//GetItemByID fetches an item from the database by its ID.
func (handler *BaseHandler) GetItemByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, _ := strconv.ParseInt(idStr, 0, 0)

	//Don't get the database.
	details, _ := models.GetItemByID(handler.db, id)
	res, err := json.Marshal(details)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.Write(res)
}

//GetItems fetches all the items from the database.
func (handler *BaseHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	allItems := models.GetItems(handler.db)
	res, err := json.Marshal(allItems)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.Write(res)
}

//UpdateItem updates an item in the database.
func (handler *BaseHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	updateItem := models.Item{}
	utils.ParseBody(r, &updateItem)

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, _ := strconv.ParseInt(idStr, 0, 0)

	itemDetails, db := models.GetItemByID(handler.db, id)

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

	w.Write(res)
}

//DeleteItem deletes an item from the database by its ID.
func (handler *BaseHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, _ := strconv.ParseInt(idStr, 0, 0)

	item := models.DeleteItem(handler.db, id)

	res, err := json.Marshal(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.Write(res)
}

//ExportItems generates and exports a CSV file with all product data.
func (handler *BaseHandler) ExportItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Disposition", "attachment; filename=data.csv")

	allItems := models.GetItems(handler.db)

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
	w.Write(buffer.Bytes())
}

//HandlePreFlightCors is a generic function to handle pre-flight requests. DELETE and PUT methods are allowed.
func (handler *BaseHandler) HandlePreFlightCors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE,PUT")
	w.WriteHeader(http.StatusNoContent)
	res, _ := json.Marshal(map[string]string{"Status": "204"})
	w.Write(res)
}
