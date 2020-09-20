package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	logger *log.Logger
}

func NewHello(logger *log.Logger) *Hello {
	return &Hello{
		logger,
	}
}

func (hello *Hello) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	hello.logger.Printf("Hi Primz...")
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "Someting wrong!!!", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(res, "Hi %s...", data)
}
