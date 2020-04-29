package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/uniplaces/carbon"
	"log"
	"net/http"
	"os"
)

func MovieHandler(writer http.ResponseWriter, request *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	logFile, err := os.OpenFile(os.Getenv("LOG_PATH")+"golang-"+carbon.Now().DateString(),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0444)
	if err != nil {
		log.Println(err)
	}
	defer logFile.Close()
	logger := log.New(logFile, "", log.LstdFlags)

	logger.Println(request.URL.Path, "is triggered.")

	parameters := mux.Vars(request)
	writer.WriteHeader(http.StatusOK)

	fmt.Fprintf(writer, "The movie category: %v\n", parameters["category"])
	fmt.Fprintf(writer, "The movie ID: %v", parameters["id"])
}
