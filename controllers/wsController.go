package controllers

import (
	"log"
	"net/http"

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

func HandleWebsocketConnection(c *gin.Context) {

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print(err.Error())
		return
	}

	defer ws.Close()

	clients[ws] = true

	for {
		var msg models.Notification

		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Print(err.Error())
			delete(clients, ws)
			break
		}
	}
}

func HandleBroadcast() {
	for {
		msg := <-broadcast

		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Print(err.Error())
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func SendNotification(data models.Notification) error {
	broadcast <- data
	return nil
}
