package controllers

import (
	"log"
	"net/http"

	"github.com/capstone-be/models"
	"github.com/capstone-be/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetCameras(c *gin.Context) {
	collection := utils.DB.Collection("Camera")

	var cameras []models.Camera

	filter := bson.M{}
	cursor, err := collection.Find(c, filter)
	if err != nil {
		c.AbortWithStatus(http.StatusBadGateway)
		log.Print(err.Error())
		return
	}

	err = cursor.All(c, &cameras)
	if err != nil {
		c.AbortWithStatus(http.StatusBadGateway)
		log.Print(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successful",
		"data":    cameras,
	})
}

func AddCameras(c *gin.Context) {
	collection := utils.DB.Collection("Camera")

	var camera models.Camera

	err := c.ShouldBindJSON(&camera)
	if err != nil {
		c.AbortWithStatus(http.StatusBadGateway)
		log.Print(err.Error())
		return
	}

	_, err = collection.InsertOne(c, camera)
	if err != nil {
		c.AbortWithStatus(http.StatusBadGateway)
		log.Print(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successful",
	})
}
