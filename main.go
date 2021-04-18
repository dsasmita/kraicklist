package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"challenge.haraj.com.sa/kraicklist/entity"
	"challenge.haraj.com.sa/kraicklist/handling"
	"challenge.haraj.com.sa/kraicklist/repository"
	"challenge.haraj.com.sa/kraicklist/service"
)

func main() {
	search := entity.Searcher{}
	recordRepository := repository.PrepareRecordRepository(search)
	recordService := service.PrepareRecordServices(recordRepository)

	// define http handlers
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/search", handling.HandleSearch(recordService))
	// define port, we need to set it as env for Heroku deployment
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}
	// start server
	fmt.Printf("Server is listening on %s...", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatalf("unable to start server due: %v", err)
	}
}
