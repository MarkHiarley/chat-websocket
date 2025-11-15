package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/markHiarley/projetinho/internal/model"
)

var broadcast = make(chan model.Message)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[*websocket.Conn]bool)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Printf("failed to upgrade to websocket: %v", err)
		return
	}

	defer ws.Close()
	clients[ws] = true

	for {
		var msg model.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error reading json: %v", err)
			delete(clients, ws)
			break
		}

		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast

		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error writing to client: %v", err)
				delete(clients, client)
				client.Close()
			}
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnections)
	go handleMessages()
	log.Printf("starting server on :9090")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
