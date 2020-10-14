package repositories

import (
	"context"

	"github.com/isurusiri/goapiboilerplate/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Create an object in the database
func Create(ctx context.Context, collection *mongo.Collection, user *models.User) error {
	_, err := collection.InsertOne(ctx, user)
	return err
}

// GetAll list all Users
func GetAll(ctx context.Context, collection *mongo.Collection) ([]*models.User, error) {
	filter := bson.D{{}}
	return filterUsers(ctx, collection, filter)
}

func filterUsers(ctx context.Context, collection *mongo.Collection, filter interface{}) ([]*models.User, error) {
	// A slice of users  for storing the decoded documents
	var users []*models.User

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return users, err
	}

	for cursor.Next(ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return users, err
		}

		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return users, err
	}

	cursor.Close(ctx)

	if len(users) == 0 {
		return users, mongo.ErrNoDocuments
	}

	return users, nil
}

// UpdateOne User instance
func UpdateOne(ctx context.Context, collection *mongo.Collection, user *models.User) (int64, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: user.ID}}
	// TODO: is this the right way to do update
	updatedUser := bson.D{{"$set", bson.D{{user.Username, user.Email}}}}

	result, err := collection.UpdateOne(ctx, filter, updatedUser)
	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil
}

// Delete a User instance.
func Delete(ctx context.Context, collection *mongo.Collection, userID primitive.ObjectID) (int64, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: userID}}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}

// FindOne returns a user identified by user id
func FindOne(ctx context.Context, collection *mongo.Collection, userID primitive.ObjectID) (models.User, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: userID}}

	var user models.User
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}
