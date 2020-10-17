package dto

// UserDto represents a User object that would be
// exposed by the API.
type UserDto struct {
	ID       string `bson:"id"`
	Username string `bson:"username"`
	Email    string `bson:"email"`
	IsActive bool   `bson:"is_active"`
}
