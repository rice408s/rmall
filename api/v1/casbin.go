package api

import (
	"github.com/gin-gonic/gin"
	"rmall/model/request"
	"rmall/service"
)

// AddPolicy 添加策略
// @Summary 添加策略
// @Description 添加策略
// @Tags 策略管理
// @Accept json
// @Produce json
// @Param request body request.AddPolicyReq
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /casbin/addPolicy [post]
func AddPolicy(c *gin.Context) {
	var req request.AddPolicyReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err = service.AddPolicy(&req)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"msg":     "添加成功",
	})
}

// RemovePolicy 删除策略
// @Summary 删除策略
// @Description 删除策略
// @Tags 策略管理
// @Accept json
// @Produce json
// @Param request body request.RemovePolicyReq
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /casbin/removePolicy [post]
func RemovePolicy(c *gin.Context) {
	var req request.RemovePolicyReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err = service.RemovePolicy(&req)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"msg":     "删除成功",
	})
}

// UpdatePolicy 修改策略
// @Summary 修改策略
// @Description 修改策略
// @Tags 策略管理
// @Accept json
// @Produce json
// @Param request body request.UpdatePolicyReq
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /casbin/updatePolicy [post]
func UpdatePolicy(c *gin.Context) {
	var req request.UpdatePolicyReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err = service.UpdatePolicy(&req)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"msg":     "修改成功",
	})
}

// GetPolicy 获取策略
// @Summary 获取策略
