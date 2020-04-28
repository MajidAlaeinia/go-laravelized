package main

import (
	"github.com/playground/app/routes"
	"log"
	"net/http"
	"time"
)

func main() {
	router := routes.WebRoutes()

	srv := &http.Server{
		Addr:              "localhost:9990",
		Handler:           router,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
