package helpers

import (
	"restaurant-management/pkg/config"
	"restaurant-management/pkg/types"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GenerateAllTokens(email, userType string) (signedToken string, signedRefreshToken string, err error) {
	claims := &types.SignedDetails{
		Email:    email,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(2)).Unix(),
		},
	}
	refreshClaims := types.SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.LocalConfig.SECRETKEY))
	if err != nil {
		return token, "", err
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(config.LocalConfig.SECRETKEY))

	if err != nil {
		return token, refreshToken, err
	}
	return token, refreshToken, nil
}

func ValidateToken(userToken string) (bool, error) {
	claims := &types.SignedDetails{}
	temp := false
	token, err := jwt.ParseWithClaims(userToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.LocalConfig.SECRETKEY), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {

			return temp, err
		}

		return temp, err
	}
	if !token.Valid {

		return temp, err
	}
	temp = true

	return temp, nil

}
