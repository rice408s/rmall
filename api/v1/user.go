package api

import (
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"rmall/model/request"
	"rmall/model/response"
	"rmall/service"
	"rmall/utils"
)

// Register @Summary 用户注册
//	@Description	用户注册
//	@Tags			用户
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.RegisterReq		true	"注册"
//	@Success		200		{object}	response.RegisterResp	"注册成功"
//	@Failure		500		{object}	string					"参数错误"
//	@Router			/user/register [post]
func Register(c *gin.Context) {

	var req request.RegisterReq
	//绑定参数
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := service.Register(&req)
	if err != nil {
		//todo 判断错误类型
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.RegisterResp{
		Id:       id,
		Username: req.Username,
		Mobile:   req.Mobile,
		Email:    req.Email,
	})

}

// Login @Summary 用户登录
//	@Description	用户登录
//	@Tags			用户
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.LoginReq	true	"登录"
//	@Success		200		{object}	response.LoginResp	"登录成功"
//	@Router			/user/login [post]
func Login(c *gin.Context) {
	var req request.LoginReq
	//绑定参数
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return

	}
	// 检查用户名是否存在
	user, err := service.FindUserByUsername(req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "用户名不存在",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return

	}
	//校验密码
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "密码错误",
		})
		return
	}

	//生成token
	token, expire, err := utils.CreateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.LoginResp{
		Token:  token,
		Expire: expire,
	})
}

// GetUserInfo @Summary 获取用户信息
//	@Description	获取用户信息
//	@Tags			用户
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string				true	"token"
//	@Success		200				{object}	response.UserInfo	"获取用户信息成功"
//	@Router			/user/info [get]
func GetUserInfo(c *gin.Context) {
	user, _ := c.Get("user")
	claims := user.(jwt.MapClaims)
	c.JSON(http.StatusOK, response.UserInfo{
		Id:       int(claims["id"].(float64)),
		Username: claims["username"].(string),
		Mobile:   claims["mobile"].(string),
		Email:    claims["email"].(string),
	})

}
