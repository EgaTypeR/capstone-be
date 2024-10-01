package controllers

import (
	"context"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/capstone-be/models"
	"github.com/capstone-be/utils"
	"github.com/gin-gonic/gin"
)

func SendCrimeEvent(c *gin.Context) {
	var request models.CrimeEvent

	// request.CameraID = c.PostForm("camera_id")
	request.CrimeType = c.PostForm("crime_type")
	request.ConfidenceScore, _ = strconv.ParseFloat(c.PostForm("confident_score"), 32)
	request.Description = c.PostForm("description")
	request.CrimeType = c.PostForm("crime_type")
	request.DetectedAt, _ = time.Parse(time.RFC3339, c.PostForm("detected_at"))

	collection := utils.DB.Collection("CrimeEvent")

	file, err := c.FormFile("footage")
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Print(err.Error())
		return
	}

	request.FootagePath = "footage_" + time.Now().UTC().Format("2006-01-02-15-04-05") + filepath.Ext(file.Filename)

	err = utils.SaveSingleFileToStorage(c, "./storage/crime-footage", "footage", request.FootagePath)
	if err != nil {
		c.AbortWithStatus(http.StatusBadGateway)
		log.Print(err.Error())
		return
	}

	response, err := collection.InsertOne(context.TODO(), request)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Print(err.Error())
		return
	}

	c.JSON(
		http.StatusAccepted,
		gin.H{
			"message": "crime event added",
			"data_id": response,
		})
}
