package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func MovieHandler(writer http.ResponseWriter, request *http.Request) {
	parameters := mux.Vars(request)
	writer.WriteHeader(http.StatusOK)

	fmt.Fprintf(writer, "The movie category: %v\n", parameters["category"])
	fmt.Fprintf(writer, "The movie ID: %v", parameters["id"])
}
