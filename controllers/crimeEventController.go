package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/capstone-be/models"
	"github.com/capstone-be/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// func SendCrimeEvent(c *gin.Context) {
// 	var request models.CrimeEvent

// 	// request.CameraID = c.PostForm("camera_id")
// 	request.CrimeType = c.PostForm("crime_type")
// 	request.ConfidenceScore, _ = strconv.ParseFloat(c.PostForm("confident_score"), 32)
// 	request.Description = c.PostForm("description")
// 	request.CrimeType = c.PostForm("crime_type")
// 	request.DetectedAt, _ = time.Parse(time.RFC3339, c.PostForm("detected_at"))

// 	collection := utils.DB.Collection("CrimeEvent")

// 	file, err := c.FormFile("footage")
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusBadRequest)
// 		log.Print(err.Error())
// 		return
// 	}

// 	request.FootagePath = "footage_" + time.Now().UTC().Format("2006-01-02-15-04-05") + filepath.Ext(file.Filename)

// 	err = utils.SaveSingleFileToStorage(c, "./storage/crime-footage", "footage", request.FootagePath)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusBadGateway)
// 		log.Print(err.Error())
// 		return
// 	}

// 	response, err := collection.InsertOne(context.TODO(), request)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusBadRequest)
// 		log.Print(err.Error())
// 		return
// 	}

// 	c.JSON(
// 		http.StatusAccepted,
// 		gin.H{
// 			"message": "crime event added",
// 			"data_id": response,
// 		})
// }

func SendCrimeEventV2(c *gin.Context) {
	var request models.RequestDetector
	var crimeEvent *models.CrimeEvent
	var notification models.Notification

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Print(err.Error())
		return
	}

	crimeEvent, err = utils.MapRequestDetectorToCrimeEvent(request)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Print(err.Error())
		return
	}

	// Simpan Ke Database
	collection := utils.DB.Collection("CrimeEvent")
	res, err := collection.InsertOne(c, crimeEvent)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}

	eventID, _ := res.InsertedID.(primitive.ObjectID)

	// Send notif
	notification.EventID = eventID
	notification.Read = false
	notification.SentAt = time.Now()
	notification.Message = "this is message"
	err = SendNotification(notification)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"message": "successful",
	})

}
