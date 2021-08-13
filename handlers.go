package main

import (
	"log"
	"net/http"
)

type InventoryHandler struct {
	l *log.Logger
}

func NewInventoryHandler(l *log.Logger) *InventoryHandler {
	return &InventoryHandler{l}
}

func (i *InventoryHandler) listInventory(w http.ResponseWriter, r *http.Request) {
	i.l.Println("Handle GET Inventory")

	// retreive from DB the inventory list. Right now now DB is implemented and data is mocked
	inventoryList := getProducts()

	err := inventoryList.toJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}
