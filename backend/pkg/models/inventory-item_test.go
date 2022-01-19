package models

import (
	"inventory-tracking/backend/pkg/utils"
	"os"
	"testing"
)

var TestBaseHandler *utils.TestBaseHandler

var testItem = Item{
	Name:     "Test sobject",
	Quantity: 5,
	Location: "Test Location",
}

func TestMain(m *testing.M) {
	TestBaseHandler = utils.CreateTestDB()

	//Run tests.
	exitVal := m.Run()

	//Close DB connection.
	// sqlDB, _ := TestBaseHandler.DB.DB()
	// sqlDB.Close()

	os.Remove("sqliteDB.db")
	os.Exit(exitVal)
}

func TestCreateItem(t *testing.T) {

	createItem := testItem.CreateItem(TestBaseHandler.DB)

	if createItem != testItem {
		t.Error("CreateItem failed.")
	}
}

func TestGetItemByID(t *testing.T) {

	getItem, _ := GetItemByID(TestBaseHandler.DB, 1)

	//Check individual items since getItem also has values like create and update dates.
	if getItem.Name != testItem.Name || getItem.Quantity != testItem.Quantity || getItem.Location != testItem.Location {
		t.Error("GetItemByID failed.")
	}
}
