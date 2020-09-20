package handlers

import (
	"golang_microservice/plant-api/data"
	"log"
	"net/http"
)

type Plants struct {
	logger *log.Logger
}

func NewPlants(logger *log.Logger) *Plants {
	return &Plants{
		logger,
	}
}

func (plants *Plants) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		plants.getPlants(res, req)
		return
	}

	res.WriteHeader(http.StatusMethodNotAllowed)
}

func (plants *Plants) getPlants(res http.ResponseWriter, req *http.Request) {
	plantsList := data.GetAllPlants()
	responseError := plantsList.ToJSON(res)

	if responseError != nil {
		http.Error(res, "JSON marshaling failed.", http.StatusInternalServerError)
	}
}
