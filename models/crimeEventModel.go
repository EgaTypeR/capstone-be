package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CrimeEvent struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	CameraID         primitive.ObjectID `bson:"camera_id"`
	CrimeType        string             `bson:"crime_type"`
	DetectedAt       time.Time          `bson:"detected_at"`
	ConfidenceScore  float64            `bson:"confidence_score"`
	Description      string             `bson:"description"`
	IsReviewed       bool               `bson:"is_reviewed"`
	FootagePath      string             `bson:"footage_path"`
	NotificationSent bool               `bson:"notification_sent"`
}
