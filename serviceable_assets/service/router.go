package service

import (
	"github.com/gorilla/mux"
)

func RouterHandler() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/demand/serviceable_assets", listServiceableAsset).Methods("GET")
	router.HandleFunc("/demand/serviceable_assets", createServiceableAsset).Methods("POST")

	return router
}
