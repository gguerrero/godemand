package service

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"../domain/models"
)

// GET /demand/address_books
// Will return all the current address_book
func listAddessBook(writer http.ResponseWriter, reader *http.Request) {
	start := time.Now()

	// randSleep()
	jsonResponse(writer, models.AddressBooks)

	log.Print("GET /demand/address_books [", time.Since(start), "]")
}

// POST /demand/address_books
// Creates and address book entry and returns it
func createAddessBook(writer http.ResponseWriter, reader *http.Request) {
	start := time.Now()

	jsonResponse(writer, "createAddressBook coming soon...")

	log.Print("POST /demand/address_books [", time.Since(start), "]")
}

func randSleep() {
	r := rand.Intn(200)
	time.Sleep(time.Duration(r) * time.Millisecond)
}
