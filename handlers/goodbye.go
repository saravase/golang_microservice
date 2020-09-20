package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	logger *log.Logger
}

func NewGoodbye(logger *log.Logger) *Goodbye {
	return &Goodbye{
		logger,
	}
}

func (goodbye *Goodbye) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	goodbye.logger.Printf("Goodbeye Primz...")
	res.Write([]byte("Goodbye Primz..."))
}
