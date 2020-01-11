package service

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"../domain/models"
)

// GET /demand/serviceable_assets
// Will return all the current address_book
func listServiceableAsset(writer http.ResponseWriter, reader *http.Request) {
	start := time.Now()

	randSleep()
	jsonResponse(writer, models.ServiceableAssets)

	log.Print("GET /demand/serviceable_assets [", time.Since(start), "]")
}

// POST /demand/serviceable_assets
// Creates and address book entry and returns it
func createServiceableAsset(writer http.ResponseWriter, reader *http.Request) {
	start := time.Now()

	jsonResponse(writer, "createAddressBook coming soon...")

	log.Print("POST /demand/serviceable_assets [", time.Since(start), "]")
}

func randSleep() {
	r := rand.Intn(1500)
	time.Sleep(time.Duration(r) * time.Millisecond)
}
