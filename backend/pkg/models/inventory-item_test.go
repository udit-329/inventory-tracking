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
	TestBaseHandler = utils.CreateTestDBHandler()

	//Run tests.
	exitVal := m.Run()

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

func TestGetItems(t *testing.T) {
	getItems := GetItems(TestBaseHandler.DB)

	//Check individual items since getItems also has values like create and update dates.
	if getItems[0].Name != testItem.Name || getItems[0].Quantity != testItem.Quantity || getItems[0].Location != testItem.Location {
		t.Error("GetItems failed.")
	} else if len(getItems) != 1 {
		t.Error("GetItems did not fetch the correct number of items..")
	}
}

func TestDeleteItems(t *testing.T) {
	deleteItem := DeleteItem(TestBaseHandler.DB, 1)

	if deleteItem.Name != testItem.Name || deleteItem.Quantity != testItem.Quantity || deleteItem.Location != testItem.Location {
		t.Error("DeleteItem failed.")
	}

	//Check that the number of items is reduced.
	getItems := GetItems(TestBaseHandler.DB)

	if len(getItems) != 0 {
		t.Error("DeleteItem failed to delete any item.")
	}

}
