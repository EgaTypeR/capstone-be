package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Camera struct for the Cameras collection
type Camera struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Num      int                `bson:"camera_num"`
	Name     string             `bson:"camera_name"`
	URL      string             `bson:"camera_url"`
	Location string             `bson:"location"`
	IsActive bool               `bson:"is_active"`
	AddedAt  time.Time          `bson:"added_at"`
}
