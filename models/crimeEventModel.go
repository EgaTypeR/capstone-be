package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Store Crime Event Data, has the same structur as the database,  has relation with Camera
type CrimeEvent struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	CameraID         primitive.ObjectID `bson:"camera_id" json:"camera_id"`
	CrimeType        CrimeType          `bson:"crime_type" json:"crime_type"`
	DetectedAt       time.Time          `bson:"detected_at" json:"detected_at"`
	ConfidenceScore  float64            `bson:"confidence_score" json:"confidence_score"`
	Description      string             `bson:"description" json:"description"`
	Danger           bool               `bson:"danger" json:"danger"`
	Dispatched       bool               `bson:"dispatched" json:"dispatched"`
	Done             bool               `bson:"done" json:"done"`
	Frame            int                `json:"frame"`
	FootagePath      string             `bson:"footage_path" json:"footage_path"`
	NotificationSent bool               `bson:"notification_sent" json:"notification_sent"`
}

// Data on History Page subset of [CrimeEvent join with Camera]
type HistoryAlert struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	CameraID   primitive.ObjectID `bson:"camera_id" json:"camera_id"`
	CameraNum  uint               `bson:"camera_num" json:"camera_num"`
	Location   string             `bson:"location"`
	DetectedAt time.Time          `bson:"detected_at" json:"detected_at"`
	Danger     bool               `bson:"danger" json:"danger"`
	Dispatched bool               `bson:"dispatched" json:"dispatched"`
	Done       bool               `bson:"done" json:"done"`
}

// Used for API update dispatched or done, subset of CrimeEvent
type UpdateAlert struct {
	Dispatched *bool `bson:"dispatched" json:"dispatched"`
	Done       *bool `bson:"done" json:"done"`
}

// Data structure for catch crime event data drom detection module
type RequestDetector struct {
	Frame     int    `json:"frame"`
	Group     int    `json:"group"`
	Persons   int    `json:"persons"`
	Celurit   int    `json:"celurit"`
	Pisau     int    `json:"pisau"`
	Pistol    int    `json:"pistol"`
	Weapons   int    `json:"wapons"`
	Anomaly   int    `json:"anomaly"`
	Status    int    `json:"status"`
	Timestamp string `json:"timestamp"`
	FileName  string `json:"file_name"`
}

type CrimeType struct {
	Group   int `json:"group"`
	Persons int `json:"persons"`
	Celurit int `json:"celurit"`
	Pisau   int `json:"pisau"`
	Pistol  int `json:"pistol"`
	Weapons int `json:"wapons"`
	Anomaly int `json:"anomaly"`
}
