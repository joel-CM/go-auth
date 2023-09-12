package services

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joel-CM/go-auth/app/models"
)

func CreateToken(user models.UserSignInModel) (string, error) {
	claims := jwt.MapClaims{"email": user.Email}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECURE_KEY")))
}
