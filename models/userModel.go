package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Username  string             `bson:"username"`
	Password  string             `bson:"password"`
	Email     string             `bson:"email"`
	Role      string             `bson:"role"`
	CreatedAt time.Time          `bson:"created_at"`
	LastLogin time.Time          `bson:"last_login"`
}
