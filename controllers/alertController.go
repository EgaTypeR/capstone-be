package controllers

import (
	"log"
	"net/http"

	"github.com/capstone-be/models"
	"github.com/capstone-be/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetHistoryAlerts(c *gin.Context) {
	filter := c.Query("filter")

	collection := utils.DB.Collection("CrimeEvent")

	pipeline := mongo.Pipeline{
		{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "Camera"},
				{Key: "localField", Value: "camera_id"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "camera_info"},
			}},
		},
		{
			{Key: "$unwind", Value: "$camera_info"},
		},
	}

	var matchStage bson.D
	switch filter {
	case "done":
		matchStage = bson.D{{Key: "done", Value: true}, {Key: "verification", Value: true}}
	case "ongoing":
		matchStage = bson.D{{Key: "done", Value: false}, {Key: "verification", Value: true}}
	case "unverif":
		matchStage = bson.D{{Key: "verification", Value: false}}
	default:
		matchStage = bson.D{{Key: "verification", Value: true}}
	}

	pipeline = append(pipeline, bson.D{{Key: "$match", Value: matchStage}})

	cursor, err := collection.Aggregate(c, pipeline)
	if err != nil {
		log.Print(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	defer cursor.Close(c)

	var result []bson.M

	err = cursor.All(c, &result)
	if err != nil {
		log.Print(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successful",
		"data":    result,
	})
}

func UpdateAlert(c *gin.Context) {
	collection := utils.DB.Collection("CrimeEvent")

	alert_id := c.Param("id")

	var updateParams models.UpdateAlert

	err := c.ShouldBindJSON(&updateParams)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Print(err.Error())
		return
	}

	objID, err := primitive.ObjectIDFromHex(alert_id)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Print(err.Error())
		return
	}

	updateField := bson.M{}
	if updateParams.Dispatched != nil {
		updateField["dispatched"] = *updateParams.Dispatched
	}
	if updateParams.Done != nil {
		updateField["done"] = *updateParams.Done
	}
	if updateParams.Verification != nil {
		updateField["verification"] = *updateParams.Verification
	}

	if len(updateField) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No fields to update"})
		return
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updateField}

	res, err := collection.UpdateOne(c, filter, update)
	if err != nil {
		c.AbortWithStatus(http.StatusBadGateway)
		log.Print(err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "alert updated",
		"result":  res,
	})

}
