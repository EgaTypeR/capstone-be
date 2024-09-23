package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Footage struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	CameraID     primitive.ObjectID `bson:"camera_id"`
	FootageStart time.Time          `bson:"footage_start"`
	FootaeEnd    time.Time          `bson:"footage_end"`
	StoragePath  string             `bson:"storage_path"`
	CreatedAt    time.Time          `bson:"created_at"`
}
