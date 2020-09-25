package handlers

import (
	"golang_microservice/plant-api/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Plants struct {
	logger *log.Logger
}

func NewPlants(logger *log.Logger) *Plants {
	return &Plants{
		logger,
	}
}

// ServeHTTP is the main entry point for the handler and statify http.Handler interface
func (plants *Plants) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	//Handle the request for a list of plants
	if req.Method == http.MethodGet {
		plants.getPlants(res, req)
		return
	}

	//Handle the request to create the new plant
	if req.Method == http.MethodPost {
		plants.createPlant(res, req)
		return
	}

	//Handle the request to update the available plant based on id
	if req.Method == http.MethodPut {

		path := req.URL.Path
		reg := `/([0-9]+)`
		regC := regexp.MustCompile(reg)
		idMatchList := regC.FindAllStringSubmatch(path, -1)

		if len(idMatchList) != 1 {
			plants.logger.Printf("Request URL iDList : %#v", idMatchList)
			http.Error(res, "Request URL have more than one id", http.StatusBadRequest)
			return
		}

		if len(idMatchList[0]) != 2 {
			plants.logger.Printf("Request URL caturedIDList : %#v", idMatchList[0])
			http.Error(res, "Regex found more id capture in the request URL", http.StatusBadRequest)
			return
		}

		idString := idMatchList[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			plants.logger.Printf("Unable to convert id to int %v", err)
			http.Error(res, "Unable to convert id to int", http.StatusBadRequest)
			return
		}

		plants.updatePlant(id, res, req)

	}

	// Handle the request to delete the plant based on id
	if req.Method == http.MethodDelete {

		path := req.URL.Path
		reg := `/([0-9]+)`
		regC := regexp.MustCompile(reg)
		idMatchList := regC.FindAllStringSubmatch(path, -1)

		if len(idMatchList) != 1 {
			plants.logger.Printf("Request URL iDList : %#v", idMatchList)
			http.Error(res, "Request URL have more than one id", http.StatusBadRequest)
			return
		}

		if len(idMatchList[0]) != 2 {
			plants.logger.Printf("Request URL caturedIDList : %#v", idMatchList[0])
			http.Error(res, "Regex found more id capture in the request URL", http.StatusBadRequest)
			return
		}

		idString := idMatchList[0][1]
		id, err := strconv.Atoi(idString)
		plants.logger.Println("ID", id)

		if err != nil {
			plants.logger.Printf("Unable to convert id to int %v", err)
			http.Error(res, "Unable to convert id to int", http.StatusBadRequest)
			return
		}

		plants.deletePlant(id, res, req)
	}

	//Catch all
	//If no method is statisfied return an error
	res.WriteHeader(http.StatusMethodNotAllowed)
}

//getPlants returns the plants from the datastore
func (plants *Plants) getPlants(res http.ResponseWriter, req *http.Request) {
	plantsList := data.GetAllPlants()
	marshalError := plantsList.ToJSON(res)

	if marshalError != nil {
		http.Error(res, "JSON marshaling failed.", http.StatusInternalServerError)
	}
}

//createPlants used to insert the new plant in the datastore
func (plants *Plants) createPlant(res http.ResponseWriter, req *http.Request) {

	plant := &data.Plant{}
	marshalError := plant.FromJSON(req.Body)

	if marshalError != nil {
		http.Error(res, "JSON Unmarshaling failed.", http.StatusBadRequest)
	}

	plants.logger.Printf("Plant : %#v", plant)
	data.AddPlant(plant)
	plantsList := &data.Plants{plant}
	plantsList.ToJSON(res)
}

//updatePlant used to update the plant data in the datastore based on id.
func (plants *Plants) updatePlant(id int, res http.ResponseWriter, req *http.Request) {

	plant := &data.Plant{}
	marshalError := plant.FromJSON(req.Body)

	if marshalError != nil {
		http.Error(res, "JSON Unmarshaling failed.", http.StatusBadRequest)
	}

	plants.logger.Printf("Plant : %#v", plant)
	err := data.UpdatePlant(id, plant)

	if err != nil {
		http.Error(res, "Plant not found.", http.StatusNotFound)
	}
	res.WriteHeader(http.StatusOK)
}

//deletePlant used to delte the plant data in the datastore based on id.
func (plants *Plants) deletePlant(id int, res http.ResponseWriter, req *http.Request) {

	err := data.DeletePlant(id)

	if err != nil {
		http.Error(res, "Plant not found.", http.StatusNotFound)
	}
	res.WriteHeader(http.StatusOK)
}
