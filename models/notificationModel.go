package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notification struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	EventID  primitive.ObjectID `bson:"event_id" json:"event_id"`
	CameraID primitive.ObjectID `bson:"camera_id" json:"camera_id"`
	UserID   primitive.ObjectID `bson:"user_id" json:"user_id"`
	Danger   bool               `bson:"danger" json:"danger"`
	Message  string             `bson:"message" json:"message"`
	Read     bool               `bson:"read" json:"read"`
	SentAt   time.Time          `bson:"sent_at" json:"sent_at"`
}

type IDs struct {
	IDs []primitive.ObjectID `json:"ids"`
}
