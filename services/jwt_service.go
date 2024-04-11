package services

import (
	"Application/models"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateToken(userId int32, name string, email string, role string) string {
	sleepTime := 12
	claims := &models.Claims{
		UserID: userId,
		Name:   name,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour * time.Duration(sleepTime))},
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, tokenStringError := token.SignedString(models.GetSecretKey())
	if tokenStringError == nil {
		return tokenString
	}
	return ""
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpectes signing method")
		}
		return models.GetSecretKey(), nil
	})
}
