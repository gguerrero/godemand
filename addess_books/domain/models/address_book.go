package models

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type AddressBook struct {
	ID         int       `json:"id"`
	CustomerID string    `json:"customerId"`
	MarketID   string    `json:"market_id"`
	IsDefault  bool      `json:"isDefault"`
	Name       string    `json:"name"`
	Note       string    `json:"note"`
	Address    Address   `json:"address"`
	Lat        string    `json:"lat"`
	Long       string    `json:"long"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type Address interface{}

var AddressBooks []AddressBook

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(pwd + "/domain/models/data/address_books.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&AddressBooks)
	if err != nil {
		log.Fatal(err)
	}
}
