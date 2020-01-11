package main

import (
	"log"
	"net/http"
	"time"

	_ "./domain/models"
	"./service"
)

const address string = "localhost:4001"

func main() {
	log.Print("Booting Server on " + address + "...")

	server := &http.Server{
		Handler: service.RouterHandler(),
		Addr:    address,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
