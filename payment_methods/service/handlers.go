package service

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"../domain/models"
)

// GET /demand/payment_methods
// Will return all the current address_book
func listPaymentMethod(writer http.ResponseWriter, reader *http.Request) {
	start := time.Now()

	randSleep()

	jsonResponse(writer, models.PaymentMethods)

	log.Print("GET /demand/payment_methods [", time.Since(start), "]")
}

// POST /demand/payment_methods
// Creates and address book entry and returns it
func createPaymentMethod(writer http.ResponseWriter, reader *http.Request) {
	start := time.Now()

	jsonResponse(writer, "createAddressBook coming soon...")

	log.Print("POST /demand/payment_methods [", time.Since(start), "]")
}

func randSleep() {
	r := rand.Intn(5000)
	time.Sleep(time.Duration(r) * time.Millisecond)
}
