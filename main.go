package main

import (
	"log"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	l := log.New(os.Stdout, "inventory-api", log.LstdFlags)

	// create mux
	router := mux.NewRouter()

	// create handler
	ih := NewInventoryHandler(l)

	router.HandleFunc("/", ih.listInventory)
}
