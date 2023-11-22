package utils

import (
	"github.com/dgrijalva/jwt-go"
	"rmall/model"
	"time"
)

func CreateToken(user *model.User) (string, int64, error) {
	expire := time.Now().Add(72 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.Id,
		"username": user.Username,
		"email":    user.Email,
		"mobile":   user.Mobile,
		"exp":      expire,
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", 0, err
	}
	return tokenString, expire, nil
}
