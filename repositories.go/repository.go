package repositories

import (
	"context"

	"github.com/isurusiri/goapiboilerplate/models"
	"go.mongodb.org/mongo-driver/mongo"
)

// Create an object in the database
func Create(ctx context.Context, collection *mongo.Collection, user *models.User) error {
	_, err := collection.InsertOne(ctx, user)
	return err
}
