package main

import (
	"log"
	"net/http"

	"github.com/markHiarley/projetinho/internal/services"
)

func main() {

	var handleConnections = services.HandleConnections
	var handleMessages = services.HandleMessages

	http.HandleFunc("/ws", handleConnections)
	go handleMessages()
	log.Printf("starting server on :9090")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
