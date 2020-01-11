package models

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type PaymentMethod struct {
	ID             string    `json:"id"`
	CustomerID     string    `json:"customer_id"`
	Type           string    `json:"type"`
	Last4Digit     string    `json:"last4digit"`
	Provider       string    `json:"provider"`
	CardToken      string    `json:"card_token"`
	CustomerToken  string    `json:"customer_token"`
	CardHolderName string    `json:"card_holder_name"`
	ExpiryMonth    string    `json:"expiry_month"`
	ExpiryYear     string    `json:"expiry_year"`
	Name           string    `json:"name"`
	IsDefault      bool      `json:"is_default"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

var PaymentMethods []PaymentMethod

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(pwd + "/domain/models/data/payment_methods.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&PaymentMethods)
	if err != nil {
		log.Fatal(err)
	}
}
