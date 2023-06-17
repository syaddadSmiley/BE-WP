package domain

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	Name                  string `json:"name"`
	Email                 string `json:"email"`
	EmailVerifiedAt       string `json:"email_verified_at"`
	PhoneNumber           string `json:"phone_number"`
	PhoneNumberVerifiedAt string `json:"phone_number_verified_at"`
	ProfilePhotoLink      string `json:"profile_photo_link"`
	ID                    string `json:"id"`
	Role                  string `json:"role"`
	jwt.RegisteredClaims
}

type JwtCustomRefreshClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}
