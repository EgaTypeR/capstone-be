package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Camera struct for the Cameras collection
type Camera struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Num      int                `bson:"camera_num" json:"camera_num"`
	Name     string             `bson:"camera_name" json:"camera_name"`
	URL      string             `bson:"camera_url" json:"camera_url"`
	Location string             `bson:"location" json:"location"`
	IsActive bool               `bson:"is_active" json:"is_active"`
	AddedAt  time.Time          `bson:"added_at" json:"added_at"`
}
