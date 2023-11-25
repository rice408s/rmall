package api

import (
	"net/http"
	"rmall/global"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthRequired token验证中间件
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(global.Config.JWT.SecretKey), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		//token是否有效
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token 无效",
			})
			c.Abort()
			return
		}
		c.Set("user", claims)
		c.Next()
	}
}
