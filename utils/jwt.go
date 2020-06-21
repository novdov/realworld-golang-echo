package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var JWTSecret = []byte("secret")

type jwtCustomClaims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(email string, id primitive.ObjectID) (string, error) {
	claims := &jwtCustomClaims{
		id.Hex(),
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(JWTSecret)
	if err != nil {
		return "", err
	}
	return t, nil
}
