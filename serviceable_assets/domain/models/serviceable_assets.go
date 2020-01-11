package models

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type ServiceableAsset struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	CustomerID string    `json:"customer_id"`
	Notes      string    `json:"notes"`
	Info       Info      `json:"info"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type Info interface{}

var ServiceableAssets []ServiceableAsset

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(pwd + "/domain/models/data/serviceable_assets.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&ServiceableAssets)
	if err != nil {
		log.Fatal(err)
	}
}
