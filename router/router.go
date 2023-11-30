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
		//用户相关
		user := v1.Group("/user")
		{
			user.POST("/register", api.UserRegister)
			user.POST("/login", api.UserLogin)
		}

		//管理员相关
		admin := v1.Group("/admin")
		{
			admin.POST("/register", api.AdminRegister)
			admin.POST("/login", api.AdminLogin)
		}

		//需要token验证的路由
		userRequired := v1.Group("/user")
		userRequired.Use(api.UserAuthRequired())
		{
			userRequired.GET("/info", api.GetUserInfo)
		}

		//需要token验证的路由
		adminRequired := v1.Group("/admin")
		adminRequired.Use(api.AdminAuthRequired(), api.CasbinMiddleware())
		{
			adminRequired.GET("/info", api.GetAdminInfo)
			//casbin相关
			adminRequired.POST("/policy/add", api.AddPolicy)
			adminRequired.POST("/policy/remove", api.RemovePolicy)
			adminRequired.POST("/policy/update", api.UpdatePolicy)
			adminRequired.GET("/policy/list", api.GetPolicyList)
		}

		role := adminRequired.Group("/role")
		{
			role.POST("/add", api.AddRole)
			role.POST("/update", api.UpdateRole)
			role.POST("/delete", api.DeleteRole)
			role.GET("/list", api.GetRoleList)
		}

	}

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
