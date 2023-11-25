package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"rmall/api/v1"
	_ "rmall/docs"
)

func Route() {
	//r := gin.New()
	//r.Use(utils.LoggerMiddleware())
	//r.Use(utils.RecoverMiddleware())
	r := gin.Default()
	//swag
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("/api/v1")
	{
		v1.POST("user/register", api.Register)
		v1.POST("user/login", api.Login)

		//需要token验证的路由
		v1.Use(api.AuthRequired())
		{
			v1.GET("user/info", api.GetUserInfo)
		}

	}

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
