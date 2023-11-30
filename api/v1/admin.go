package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"rmall/model/request"
	"rmall/model/response"
	"rmall/service"
)

// AdminRegister @Summary 管理员注册
//	@summary		管理员注册
//	@Description	管理员注册
//	@Tags			管理员
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.AdminRegisterReq	true	"注册"
//	@Success		200		{object}	response.AdminRegisterResp
//	@Failure		500		{object}	string
//	@Router			/admin/register [post]
func AdminRegister(c *gin.Context) {
	var req request.AdminRegisterReq
	//绑定参数
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := service.AdminRegister(&req)
	if err != nil {
		//todo 判断错误类型
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       id,
		"username": req.Username,
		"mobile":   req.Mobile,
		"email":    req.Email,
	})
}

// AdminLogin @Summary 管理员登录
//	@summary		管理员登录
//	@Description	管理员登录
//	@Tags			管理员
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.AdminLoginReq	true	"登录"
//	@Success		200		{object}	response.AdminLoginResp
//	@Router			/admin/login [post]
func AdminLogin(c *gin.Context) {
	var req request.AdminLoginReq
	//绑定参数
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, expire, err := service.AdminLogin(&req)
	if err != nil {
		//todo 判断错误类型
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":  token,
		"expire": expire,
	})
}

// GetAdminInfo @Summary 管理员信息
//	@summary		管理员信息
//	@Description	管理员信息
//	@Tags			管理员
//	@Accept			json
//	@Produce		json
//	@Param			token	header		string	true	"token"
//	@Success		200		{object}	response.AdminInfo
//	@Router			/admin/info [get]
func GetAdminInfo(c *gin.Context) {
	claims, _ := c.Get("admin")
	adminClaims := claims.(jwt.MapClaims)

	c.JSON(http.StatusOK, response.AdminInfo{
		Id:       int(adminClaims["id"].(float64)),
		Username: adminClaims["username"].(string),
		Mobile:   adminClaims["mobile"].(string),
		Email:    adminClaims["email"].(string),
	})
}
