package utils

import (
	"os"
	"restaurant-management/pkg/types"

	jwt "github.com/dgrijalva/jwt-go"
)

var SecretKey string = os.Getenv("SECRET_KEY")

func ParseToken(usertoken string) (*types.SignedDetails, error) {
	var claims = &types.SignedDetails{}
	_, err := jwt.ParseWithClaims(usertoken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	return claims, err
}

