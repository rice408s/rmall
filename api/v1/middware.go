package api

import (
	"net/http"
	"rmall/global"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// UserAuthRequired token验证中间件
func UserAuthRequired() gin.HandlerFunc {
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

func AdminAuthRequired() gin.HandlerFunc {
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
		c.Set("admin", claims)
		c.Next()
	}
}

func CasbinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求的URI和方法
		obj := c.Request.URL.RequestURI()
		act := c.Request.Method

		// 获取用户的角色
		claims, _ := c.Get("admin")
		userClaims := claims.(jwt.MapClaims)
		//获取用户名
		username := userClaims["username"].(string)
		//获取用户角色
		roles, err := global.E.GetRolesForUser(username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error occurred while getting the roles for the user",
			})
			c.Abort()
			return
		}
		// 检查每个角色的策略
		allowed := false
		for _, role := range roles {
			ok, err := global.E.Enforce(role, obj, act)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "error occurred while checking the policy",
				})
				c.Abort()
				return
			}
			if ok {
				allowed = true
				break
			}
		}

		// 如果策略不允许，返回403错误
		if !allowed {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "forbidden",
			})
			c.Abort()
			return
		}

		// 如果策略允许，继续处理请求
		c.Next()
	}
}
