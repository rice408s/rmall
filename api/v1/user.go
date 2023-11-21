package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rmall/model/request"
	"rmall/service"
)

// Register @Summary 用户注册
// @Description 用户注册
// @Tags 用户
// @Accept json
// @Produce json
// @Param request body request.RegisterReq true "注册"
// @Success 200 {object} string "注册成功"
// @Router /api/v1/user/register [post]
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})

}
