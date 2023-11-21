package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "rmall/docs"
	"rmall/initialize"
)

// 添加注释以描述 server 信息
//	@title			swagger Example API
//	@version		1.0
//	@description	这是一个描述.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/

//	@securityDefinitions.basic	BasicAuth

//	func testLogin() {
//		loginReq := request.LoginReq{
//			Username: "whw",
//			Password: "123456",
//		}
//		var u *service.UserService
//		login, i, err := u.Login(&loginReq)
//		fmt.Println(login, i, err)
//	}
func main() {
	initialize.Run()
	//testLogin()
	//err := service.Register(&request.RegisterReq{
	//	Username: "rice",
	//	Password: "123456",
	//	Mobile:   "18582622443",
	//	Email:    "rice408s@gmail.com",
	//})
	//if err != nil {
	//	fmt.Println(err)
	//}

	//DBtest()
	//r := router.PingRouter()
	//r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//v1 := r.Group("/api/v1")
	//{
	//	{
	//		v1.GET("", func(c *gin.Context) {
	//			c.JSON(200, gin.H{
	//				"message": "pong",
	//			})
	//		})
	//
	//		v1.GET("helloworld", Helloworld)
	//		v1.POST("helloName", HelloName)
	//	}
	//}
	//err := r.Run(":8080")
	//if err != nil {
	//	return
	//} //...
}

// @Summary		测试
// @Description	测试
// @Tags			测试
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"ID"
// @Success		200	{object}	test
// @Router			/api/v1/test/{id} [get]
type test struct {
	Id int
	D1 string
}

//	@BasePath	/

// Helloworld PingExample godoc
//
//	@Summary		Hello,world
//	@Schemes		json
//	@Description	测试输出hello,world
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	Helloworld
//	@Router			/api/v1/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}
