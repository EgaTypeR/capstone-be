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
	var notification models.Notification // Initialize notification directly as a struct, not a pointer

	// Parse JSON request body
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Print(err.Error())
		return
	}

	// Map request to crime event
	crimeEvent, err = utils.MapRequestDetectorToCrimeEvent(request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Print(err.Error())
		return
	}

	// Save CrimeEvent to the database
	collection := utils.DB.Collection("CrimeEvent")
	res, err := collection.InsertOne(c, crimeEvent)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to insert crime event"})
		log.Print(err.Error())
		return
	}

	// Retrieve the inserted CrimeEvent ID
	eventID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get inserted crime event ID"})
		return
	}

	// Initialize notification with necessary fields
	notification = models.Notification{
		EventID: eventID,
		Read:    false,
		SentAt:  time.Now(),
		Message: "this is message",
		Danger:  crimeEvent.Danger,
	}

	// Save Notification to the database
	notifCollection := utils.DB.Collection("Notification")
	notifRes, err := notifCollection.InsertOne(c, &notification)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to insert notification"})
		log.Print(err.Error())
		return
	}

	// Retrieve the inserted Notification ID
	notifID, ok := notifRes.InsertedID.(primitive.ObjectID)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get inserted notification ID"})
		return
	}

	notification.ID = notifID

	// Send Notification (external function)
	err = SendNotification(notification)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to send notification"})
		log.Print(err.Error())
		return
	}

	// Send JSON response
	c.JSON(http.StatusAccepted, gin.H{
		"message": "successful",
	})
}
