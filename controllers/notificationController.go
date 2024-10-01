package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/capstone-be/utils"
	"github.com/gin-gonic/gin"
)

func GetNotification(c *gin.Context) {
	collection := utils.DB.Collection("Notification")

	_, err := collection.Find(context.TODO(), "")
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Print(err.Error())
		return
	}

}
