package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
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

func GetUserIDFromJWT(c echo.Context) primitive.ObjectID {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	idStr, ok := claims["id"].(string)
	if !ok {
		return primitive.NilObjectID
	}
	id, _ := primitive.ObjectIDFromHex(idStr)
	return id
}
