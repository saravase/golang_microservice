package main

import (
	"context"
	"golang_microservice/plant-api/handlers"

	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	logger := log.New(os.Stdout, "product-plant-api", log.LstdFlags)
	plantsHandler := handlers.NewPlants(logger)

	//Create Own serveMux
	serveMux := http.NewServeMux()
	serveMux.Handle("/", plantsHandler)

	server := http.Server{
		Addr:         ":9090",
		Handler:      serveMux,
		IdleTimeout:  100 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	sig := <-signalChannel

	logger.Println("Received terminate signal, graceful timeout", sig)


	terminateContext, contextErr := context.WithTimeout(context.Background(), 30*time.Second)
	if contextErr != nil {
		logger.Fatal(contextErr)
	}

	server.Shutdown(terminateContext)

}
