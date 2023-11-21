package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"rmall/api/v1"
	_ "rmall/docs"
)

func Route() {
	r := gin.Default()
	//swag
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("/api/v1")
	{
		v1.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		v1.POST("user/register", api.Register)
	}

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
