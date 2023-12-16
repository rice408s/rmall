package middleware

import (
	"fmt"
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

// AdminAuthRequired token验证中间件
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

// CasbinMiddleware 权限中间件
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

// Cors 跨域中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, content-type, Content-Length, X-CSRF-Token, Token,session,Access-Control-Allow-Headers,account")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("Content-Type", "application/json")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("Panic info is: %v\n", err)
			}
		}()
		c.Next()
	}
}
