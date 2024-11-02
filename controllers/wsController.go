package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/capstone-be/models"
	"github.com/capstone-be/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	collection := utils.DB.Collection("Notification")
	res, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}

	notifID, _ := res.InsertedID.(primitive.ObjectID)
	data.ID = notifID
	broadcast <- data
	return nil
}
