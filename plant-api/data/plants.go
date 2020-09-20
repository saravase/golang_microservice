package data

import (
	"encoding/json"
	"io"
	"time"
)

type Plant struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float32 `json:"price"`
	CreatedAt   string  `json:"-"`
	UpdatedAt   string  `json:"-"`
	DeletedAt   string  `json:"-"`
}

type Plants []*Plant

func (plants *Plants) ToJSON(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(plants)
}

func GetAllPlants() Plants {
	return plantsList
}

var plantsList = []*Plant{
	&Plant{
		ID:          1,
		Name:        "Rose",
		Description: "Beautiful Flower",
		Category:    "Flower",
		Price:       100.00,
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
		DeletedAt:   time.Now().UTC().String(),
	},
	&Plant{
		ID:          2,
		Name:        "Apple",
		Description: "Tasty Fruit",
		Category:    "Fruit",
		Price:       500.00,
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
		DeletedAt:   time.Now().UTC().String(),
	},
}
