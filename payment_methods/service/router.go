package service

import (
	"github.com/gorilla/mux"
)

func RouterHandler() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/demand/payment_methods", listPaymentMethod).Methods("GET")
	router.HandleFunc("/demand/payment_methods", createPaymentMethod).Methods("POST")

	return router
}
