package main

import (
	"context"
	"os/signal"

	"github.com/gorilla/mux"

	"github.com/saravase/golang_microservice/plant-api/handlers"

	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	// New creates a new plant-api Logger.
	logger := log.New(os.Stdout, "product-plant-api", log.LstdFlags)

	// Initialize the plant struct properties
	plantHandler := handlers.NewPlant(logger)

	// NewRouter returns a new gorilla mux router instance
	gorillaMux := mux.NewRouter()

	/*
		Subrouter creates a subrouter for the route
		It will test the inner routes only if the parent route matched
	*/

	// Get subrouter
	getRouter := gorillaMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/plant", plantHandler.GetPlants)

	// Post subrouter
	postRouter := gorillaMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/plant", plantHandler.CreatePlant)
	postRouter.Use(plantHandler.PlantValidationMiddleware)

	// Put subrouter
	putRouter := gorillaMux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/plant/{id:[0-9]+}", plantHandler.UpdatePlant)
	putRouter.Use(plantHandler.PlantValidationMiddleware)

	// Delete subrouter
	deleteRouter := gorillaMux.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/plant/{id:[0-9]+}", plantHandler.DeletePlant)

	// Initialize the plant-api server properties
	server := http.Server{
		Addr:         ":9090",
		Handler:      gorillaMux,
		IdleTimeout:  100 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// Initialize the go-routine function
	go func() {

		// ListenAndServe listens on the TCP network address specified in the server property
		listenAndServeError := server.ListenAndServe()

		if listenAndServeError != nil {
			logger.Fatal(listenAndServeError)
		}
	}()

	// Make the channel with type os.Signal
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	// Read the channel value
	sig := <-signalChannel

	logger.Println("Received os signal, graceful timeout", sig)

	//Canceling this context releases resources associated with it
	terminateContext, terminateContextError := context.WithTimeout(context.Background(), 30*time.Second)

	if terminateContextError != nil {
		logger.Fatal(terminateContextError)
	}

	// Shutdown gracefully shuts down the server without interrupting any active connections
	server.Shutdown(terminateContext)

}
