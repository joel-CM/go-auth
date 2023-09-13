package services

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joel-CM/go-auth/app/models"
)

// create/sign token
func CreateToken(user models.UserSignInModel) (string, error) {
	claims := jwt.MapClaims{"email": user.Email}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECURE_KEY")))
}

// verify token
func VerifyToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token")
		}
		return []byte(os.Getenv("JWT_SECURE_KEY")), nil
	})

	if err != nil {
		return false, err
	}

	return token.Valid, nil
}
