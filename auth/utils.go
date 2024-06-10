package auth

import (
	"fmt"
	"gymbro-api/config"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userID uint64) (string, error) {
	cfg := config.LoadConfig()
	var secret = []byte(cfg.JWTSecret)

	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 3).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret)
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	cfg := config.LoadConfig()
	var secret = []byte(cfg.JWTSecret)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}

		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func GetParsedUserId(userId string) (uint64, error) {
	return strconv.ParseUint(userId, 10, 64)
}
