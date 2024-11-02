package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/capstone-be/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetNotification(c *gin.Context) {
	collection := utils.DB.Collection("Notification")

	filter := bson.M{}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Print(err.Error())
		return
	}

	var notification []bson.M

	if err := cursor.All(context.TODO(), &notification); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Print(err.Error())
		return
	}

	if err := cursor.Close(context.TODO()); err != nil {
		log.Printf(err.Error())
	}

	c.JSON(http.StatusOK, notification)
}
