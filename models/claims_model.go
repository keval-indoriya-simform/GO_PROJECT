package models

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
)

type Claims struct {
	UserID int32  `json:"user_id,omitempty"`
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
	Role   string `json:"role,omitempty"`
	jwt.RegisteredClaims
}

func GetSecretKey() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "jwt_secret"
	}
	return []byte(secret)
}
