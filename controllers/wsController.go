package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/capstone-be/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan models.Notification)
var pingInterval = 30 * time.Second // Set the interval for pinging the client

// HandleWebsocketConnection establishes a WebSocket connection and handles incoming messages.
func HandleWebsocketConnection(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}
	defer ws.Close()

	clients[ws] = true
	log.Println("New client connected")

	// Send a ping every pingInterval to keep the connection alive
	go sendPing(ws)

	// Read messages from the WebSocket
	for {
		var msg models.Notification
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			delete(clients, ws)
			break
		}
		// Optionally handle incoming messages here (e.g., store, validate)
	}
}

// sendPing sends a ping message to the WebSocket client at regular intervals.
func sendPing(ws *websocket.Conn) {
	ticker := time.NewTicker(pingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				log.Printf("Error sending ping: %v", err)
				return
			}
		}
	}
}

// HandleBroadcast sends notifications to all connected clients.
func HandleBroadcast() {
	for {
		msg := <-broadcast

		// Broadcast the message to all connected clients
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Error writing to client: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

// SendNotification sends a notification to all clients via the broadcast channel.
func SendNotification(data models.Notification) {
	broadcast <- data
}
