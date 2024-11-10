package utils

import (
	"fmt"
	"time"

	"github.com/capstone-be/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MapRequestDetectorToCrimeEvent(req models.RequestDetector) (*models.CrimeEvent, error) {
	var res models.CrimeEvent
	var err error

	res.CameraID, err = primitive.ObjectIDFromHex(req.CameraID.Hex())
	if err != nil {
		return nil, err
	}

	res.CrimeType.Group = req.Group
	res.CrimeType.Persons = req.Persons
	res.CrimeType.Celurit = req.Celurit
	res.CrimeType.Pisau = req.Pisau
	res.CrimeType.Weapons = req.Weapons
	res.CrimeType.Anomaly = req.Anomaly
	res.FootagePath = FootageFileName(req.FileName)
	res.Frame = req.Frame
	res.CameraID = req.CameraID

	res.DetectedAt = customParseTime(req.Timestamp)
	if req.Status == 2 {
		res.Danger = true
	} else {
		res.Danger = false
	}

	// default value
	res.Dispatched = false
	res.Done = false
	res.Verification = false
	return &res, nil
}

func MapCrimeEventToNotification(req models.CrimeEvent) models.Notification {
	var res models.Notification

	return res
}

func customParseTime(s string) time.Time {
	layout := "2006-01-02T15:04:05.000000"
	t, err := time.Parse(layout, s)
	if err != nil {
		fmt.Println(err)
	}
	return t
}
