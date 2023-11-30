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

		//需要token验证的user路由
		userRequired := v1.Group("/user")
		userRequired.Use(api.UserAuthRequired())
		{
			userRequired.GET("/info", api.GetUserInfo)
		}

		//需要token验证的admin路由
		adminRequired := v1.Group("/admin")
		adminRequired.Use(api.AdminAuthRequired(), api.CasbinMiddleware())
		{
			adminRequired.GET("/info", api.GetAdminInfo) //管理员信息
			adminRequired.GET("/list", api.GetAdminList) //管理员列表
		}

		//casbin相关路由
		policy := adminRequired.Group("/policy")
		{
			policy.POST("/add", api.AddPolicy)              //添加策略
			policy.POST("/remove", api.RemovePolicy)        //删除策略
			policy.POST("/update", api.UpdatePolicy)        //更新策略
			policy.GET("/list", api.GetPolicyList)          //获取策略列表
			policy.POST("/listByRole", api.GetPolicyByRole) //获取角色策略列表
		}

		//角色相关路由
		role := adminRequired.Group("/role")
		{
			role.POST("/add", api.AddRole)       //添加角色
			role.POST("/update", api.UpdateRole) //更新角色
			role.POST("/delete", api.DeleteRole) //删除角色
			role.GET("/list", api.GetRoleList)   //获取角色列表
			//给管理员分配角色
			role.POST("/assign", api.AssignRoleToAdmin)       //给管理员分配角色
			role.POST("/listByAdmin", api.GetRoleListByAdmin) //获取管理员角色列表
		}

	}

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
