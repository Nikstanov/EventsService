package utills

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

const (
	SECRET_KEY = "SECRET_KEY"
)

var (
	secretKey string
)

func InitJWT() {
	var exists bool
	secretKey, exists = os.LookupEnv(SECRET_KEY)
	if !exists {
		panic("The secret key is not set")
	}
}

func GenerateToken(email string, userId int64) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signed method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, errors.New("could not parse token")
	}
	if !parsedToken.Valid {
		return 0, errors.New("invalid token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	userId := claims["userId"].(float64)

	return int(userId), nil
}
