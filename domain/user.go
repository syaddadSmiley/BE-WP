package domain

import (
	"context"
)

const (
	CollectionUser = "users"
)

type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	// Fetch(c context.Context) ([]User, error)
	GetByEmail(c context.Context, email string) (User, error)
	// GetByID(c context.Context, id string) (User, error)
}
