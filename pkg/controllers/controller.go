package controllers

import (
	"encoding/json"
	"inventory-tracking/pkg/models"
	"inventory-tracking/pkg/utils"
	"net/http"
)

var newItem models.Item

func AddItem(w http.ResponseWriter, r *http.Request) {
	createItem := &models.Item{}
	utils.ParseBody(r, createItem)
	b := createItem.CreateItem()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
