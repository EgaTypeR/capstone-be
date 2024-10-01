package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CrimeEvent struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	CameraID         primitive.ObjectID `bson:"camera_id" json:"camera_id"`
	CrimeType        string             `bson:"crime_type" json:"crime_type"`
	DetectedAt       time.Time          `bson:"detected_at" json:"detected_at"`
	ConfidenceScore  float64            `bson:"confidence_score" json:"confidence_score"`
	Description      string             `bson:"description" json:"description"`
	IsReviewed       bool               `bson:"is_reviewed" json:"is_reviewed"`
	FootagePath      string             `bson:"footage_path" json:"footage_path"`
	NotificationSent bool               `bson:"notification_sent" json:"notification_sent"`
}
