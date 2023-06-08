package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `validate:"required"`
	Username string    `validate:"required"`
	Email    string    `validate:"required,email"`
}

func main() {
	// Generate a UUID
	id := uuid.New()

	// Create a new user
	user := User{
		ID:       id,
		Username: "john123",
		Email:    "john@example.com",
	}

	// Validate the user struct
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		fmt.Println("Validation error:", err)
		return
	}

	// Print the user details
	fmt.Println("User ID:", user.ID)
	fmt.Println("Username:", user.Username)
	fmt.Println("Email:", user.Email)
}
