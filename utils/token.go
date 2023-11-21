package utils

import (
	"rmall/model"
	"time"
	"github.com/dgrijalva/jwt-go"
)

func CreateToken(user *model.User) (string, int64, error) {
	expire := time.Now().Add(72 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
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


func ParseToken(tokenString string) (*jwt.Token, *jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, nil, err
	}
	return token, &claims, nil
}