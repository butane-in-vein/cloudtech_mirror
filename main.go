package main

import (
	"assignment1/exchange"
	"assignment1/info"
	"assignment1/status"
	"log"
	"net/http"
	"os"
)

func main() {
	// Pirate kendrick lamar be like -> "my port side just went viral"
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := http.NewServeMux()
	router.HandleFunc("/countryinfo/v1/status/", http.HandlerFunc(status.GetStatus))
	router.HandleFunc("/countryinfo/v1/info/{code}", http.HandlerFunc(info.GetInfo))
	router.HandleFunc("/countryinfo/v1/exchange/{code}", http.HandlerFunc(exchange.GetExchange))

	// start service
	log.Println("Starting server @ port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
