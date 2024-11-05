package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/capstone-be/models"
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

	var notification []models.Notification

	if err := cursor.All(context.TODO(), &notification); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Print(err.Error())
		return
	}

	if err := cursor.Close(context.TODO()); err != nil {
		log.Print(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successful",
		"data":    notification,
	})
}

func CountUnreadNotification(c *gin.Context) {
	collection := utils.DB.Collection("Notification")
	filter := bson.M{"read": false}

	res, err := collection.CountDocuments(c, filter)
	if err != nil {
		c.AbortWithStatus(http.StatusBadGateway)
		log.Print(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully",
		"data":    res,
	})
}

func ReadNotification(c *gin.Context) {
	collection := utils.DB.Collection("Notification")
	var data models.IDs
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Print(err.Error())
		return
	}

	filter := bson.M{"_id": bson.M{"$in": data.IDs}}
	update := bson.M{"$set": bson.M{"read": true}}

	log.Print(filter, "\n", update)
	_, err = collection.UpdateMany(c, filter, update)
	if err != nil {
		c.AbortWithStatus(http.StatusBadGateway)
		log.Print(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully",
	})

}
