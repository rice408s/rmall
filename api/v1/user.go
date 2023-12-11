package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"rmall/model/request"
	"rmall/model/response"
	"rmall/service"
)

// UserRegister @Summary 用户注册
//
//	@summary		用户注册
//	@Description	用户注册
//	@Tags			用户
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.UserRegisterReq		true	"注册"
//	@Success		200		{object}	response.UserRegisterResp	"注册成功"
//	@Failure		500		{object}	string						"参数错误"
//	@Router			/user/register [post]
func UserRegister(c *gin.Context) {

	var req request.UserRegisterReq
	//绑定参数
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := service.UserRegister(&req)
	if err != nil {
		//todo 判断错误类型
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.UserRegisterResp{
		Id:       id,
		Username: req.Username,
		Mobile:   req.Mobile,
		Email:    req.Email,
	})

}

// UserLogin @Summary 用户登录
//
//	@summary		用户登录
//	@Description	用户登录
//	@Tags			用户
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.UserLoginReq	true	"登录"
//	@Success		200		{object}	response.UserLoginResp	"登录成功"
//	@Router			/user/login [post]
func UserLogin(c *gin.Context) {
	var req request.UserLoginReq
	//绑定参数
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return

	}
	//校验用户名和密码
	token, expire, err := service.UserLogin(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.UserLoginResp{
		Token:  token,
		Expire: expire,
	})
}

// GetUserInfo @Summary 获取用户信息
//
//	@Description	获取用户信息
//	@Tags			用户
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string				true	"token"
//	@Success		200				{object}	response.UserInfo	"获取用户信息成功"
//	@Router			/user/info [get]
func GetUserInfo(c *gin.Context) {
	claims, _ := c.Get("user")
	userClaims := claims.(jwt.MapClaims)

	c.JSON(http.StatusOK, response.UserInfo{
		Id:       int(userClaims["id"].(float64)),
		Username: userClaims["username"].(string),
		Mobile:   userClaims["mobile"].(string),
		Email:    userClaims["email"].(string),
	})
}
