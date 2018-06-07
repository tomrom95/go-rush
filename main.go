package main

import (
	"log"
	"net/http"

	"github.com/tomrom95/go-rush/api"
	"github.com/tomrom95/go-rush/api/handlers"
	"github.com/tomrom95/go-rush/datastore"
)

func main() {
	store := datastore.NewDataStore()
	err := store.Connect()
	if err != nil {
		log.Fatalf("Could not connect to db: %s", err)
	}
	defer store.Close()
	router := api.NewRouter(&handlers.Context{store.DB})
	log.Fatal(http.ListenAndServe(":3333", router))
}
