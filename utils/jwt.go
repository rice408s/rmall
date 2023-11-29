package utils

import (
	"github.com/dgrijalva/jwt-go"
	"rmall/global"
	"rmall/model"
	"time"
)

// 修改成泛型
//func CreateUserToken(user interface{}) (string, int64, error) {
//	expire := time.Now().Add(time.Duration(global.Config.JWT.Expire) * time.Second).Unix()
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
//		"user": user,
//		"exp":  expire,
//	})
//
//	// 使用密钥对完整编码的令牌进行签名，并将其作为字符串获取
//	tokenString, err := token.SignedString([]byte(global.Config.JWT.SecretKey))
//	if err != nil {
//		return "", 0, err
//	}
//	return tokenString, expire, nil
//}

func CreateUserToken(user *model.User) (string, int64, error) {
	expire := time.Now().Add(time.Duration(global.Config.JWT.Expire) * time.Second).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.Id,
		"username": user.Username,
		"email":    user.Email,
		"mobile":   user.Mobile,
		"exp":      expire,
	})

	// 使用密钥对完整编码的令牌进行签名，并将其作为字符串获取
	tokenString, err := token.SignedString([]byte(global.Config.JWT.SecretKey))
	if err != nil {
		return "", 0, err
	}
	return tokenString, expire, nil
}

func CreateAdminToken(admin *model.Admin) (string, int64, error) {
	expire := time.Now().Add(time.Duration(global.Config.JWT.Expire) * time.Second).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       admin.Id,
		"username": admin.Username,
		"email":    admin.Email,
		"mobile":   admin.Mobile,
		"exp":      expire,
	})

	// 使用密钥对完整编码的令牌进行签名，并将其作为字符串获取
	tokenString, err := token.SignedString([]byte(global.Config.JWT.SecretKey))
	if err != nil {
		return "", 0, err
	}
	return tokenString, expire, nil
}
