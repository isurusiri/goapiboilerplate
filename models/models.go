package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User data model.
type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Username  string             `bson:"username"`
	email     string             `bson:"email"`
	isActive  bool               `bson:"is_active"`
}
