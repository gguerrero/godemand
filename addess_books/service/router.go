package service

import (
	"github.com/gorilla/mux"
)

func RouterHandler() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/demand/address_books", listAddessBook).Methods("GET")
	router.HandleFunc("/demand/address_books", createAddessBook).Methods("POST")

	return router
}
