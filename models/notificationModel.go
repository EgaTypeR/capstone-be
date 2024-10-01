package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notification struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	EventID          primitive.ObjectID `bson:"event_id"`
	UserID           primitive.ObjectID `bson:"user_id"`
	NotificationType string             `bson:"notification_type"`
	Message          string             `bson:"message"`
	Read             bool               `bson:"read"`
	SentAt           time.Time          `bson:"sent_at"`
}
