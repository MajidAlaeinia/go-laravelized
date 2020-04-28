package controllers

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

func ArticleHandler(writer http.ResponseWriter, request *http.Request) {
	parameters := mux.Vars(request)
	writer.WriteHeader(http.StatusOK)

	fmt.Fprintf(writer, "The article category: %v\n", parameters["category"])
	fmt.Fprintf(writer, "The book ID: %v", parameters["id"])
}
