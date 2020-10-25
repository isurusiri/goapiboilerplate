package services

import (
	"github.com/isurusiri/goapiboilerplate/dto"
	"github.com/isurusiri/goapiboilerplate/models"
)

// Create a new instance of the dto being passed.
func Create(newUser dto.UserDto) {
	user := models.User{
		Username: newUser.Username,
		Email:    newUser.Email,
		IsActive: newUser.IsActive,
	}
}

// GetAll returns all instances of an entity type.
func GetAll() {}

// GetOne returns an instance identified by the the id provided.
func GetOne() {}

// Update an instance with the values provided by the dto and
// identified by the provided id.
func Update() {}

// Delete an instance identified by the provided id.
func Delete() {}
