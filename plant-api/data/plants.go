package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Plant struct represent the plant  details
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

// Initialize the plants list
type Plants []*Plant

// Initialize Excceptions
var PlantNotFoundException = fmt.Errorf("Plant not found")

// Serialize the data into JSON
func (plants *Plants) ToJSON(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(plants)
}

// Deserialize the JSON into data
func (plant *Plant) FromJSON(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(plant)
}

// GetAllPlants is used to get the plants list
func GetAllPlants() Plants {
	return plantsList
}

// AddPlant is used to add the plant in datastore
func AddPlant(plant *Plant) {
	plant.ID = generatePlantNextID()
	plantsList = append(plantsList, plant)
}

// generatePlantNextID is used to generate the id for the newly added plant
func generatePlantNextID() int {
	plant := plantsList[len(plantsList)-1]
	return plant.ID + 1
}

func UpdatePlant(id int, plant *Plant) error {
	updatePlant, position, err := getPlantByID(id)

	if err != nil {
		return PlantNotFoundException
	}

	plant.ID = updatePlant.ID
	plantsList[position] = plant
	return nil
}

func DeletePlant(id int) error {
	_, position, err := getPlantByID(id)

	if err != nil {
		return PlantNotFoundException
	}
	plantsList = append(plantsList[:position], plantsList[position+1:]...)

	return nil
}

func getPlantByID(id int) (*Plant, int, error) {
	for position, plantData := range plantsList {
		if plantData.ID == id {
			return plantData, position, nil
		}
	}
	return nil, -1, PlantNotFoundException
}

// PlantsList datastore
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
