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

	// Swagger文档路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")

	// 用户相关路由
	user := v1.Group("/user")
	{
		user.POST("/register", api.UserRegister)
		user.POST("/login", api.UserLogin)
	}

	// 管理员相关路由
	admin := v1.Group("/admin")
	{
		admin.POST("/register", api.AdminRegister)
		admin.POST("/login", api.AdminLogin)
	}

	// 需要token验证的用户路由
	userRequired := v1.Group("/user")
	userRequired.Use(api.UserAuthRequired())
	{
		userRequired.POST("/info", api.GetUserInfo)

		// 用户订单相关路由
		order := userRequired.Group("/order")
		{
			order.POST("/create", api.CreateOrder)
			order.POST("/update", api.UpdateOrder)
			order.POST("/get", api.GetOrderById)
			//order.POST("/list", api.GetOrderList)
			order.POST("getByUid", api.GetOrderByUid)

		}
	}

	// 需要token验证的管理员路由
	adminRequired := v1.Group("/admin")
	adminRequired.Use(api.AdminAuthRequired(), api.CasbinMiddleware())
	{
		adminRequired.POST("/info", api.GetAdminInfo) // 管理员信息
		adminRequired.POST("/list", api.GetAdminList) // 管理员列表

		// Casbin相关路由
		policy := adminRequired.Group("/policy")
		{
			policy.POST("/add", api.AddPolicy)              // 添加策略
			policy.POST("/remove", api.RemovePolicy)        // 删除策略
			policy.POST("/update", api.UpdatePolicy)        // 更新策略
			policy.POST("/list", api.GetPolicyList)         // 获取策略列表
			policy.POST("/listByRole", api.GetPolicyByRole) // 获取角色策略列表
		}

		// 角色相关路由
		role := adminRequired.Group("/role")
		{
			role.POST("/add", api.AddRole)       // 添加角色
			role.POST("/update", api.UpdateRole) // 更新角色
			role.POST("/delete", api.DeleteRole) // 删除角色
			role.POST("/list", api.GetRoleList)  // 获取角色列表

			// 给管理员分配角色
			role.POST("/assign", api.AssignRoleToAdmin)       // 给管理员分配角色
			role.POST("/listByAdmin", api.GetRoleListByAdmin) // 获取管理员角色列表
		}

		// 产品管理员相关路由
		product := adminRequired.Group("/product")
		{
			product.POST("/add", api.AddProduct)       // 添加产品
			product.POST("/update", api.UpdateProduct) // 更新产品
			product.POST("/delete", api.DeleteProduct) // 删除产品
		}
		//订单管理员相关路由
		order := adminRequired.Group("/order")
		{
			order.POST("/list", api.GetOrderList)       // 获取订单列表
			order.POST("/get", api.GetOrderById)        // 获取订单详情
			order.POST("/update", api.UpdateOrder)      // 更新订单
			order.POST("/delete", api.DeleteOrder)      // 删除订单
			order.POST("/listByUid", api.GetOrderByUid) // 获取用户订单列表
			order.POST("/listByPid", api.GetOrderByPid) // 获取产品订单列表
		}
	}

	// 产品公共路由
	productCommon := v1.Group("/product")
	{
		productCommon.POST("/list", api.GetProductList)     // 获取产品列表
		productCommon.POST("/detail", api.GetProductById)   // 获取产品详情
		productCommon.POST("/search", api.GetProductByName) // 模糊查询产品
	}

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
