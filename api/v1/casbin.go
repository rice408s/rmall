package api

import (
	"github.com/gin-gonic/gin"
	"rmall/model/request"
	"rmall/service"
)

// AddPolicy 添加策略
//
//	@Summary		添加策略
//	@Description	添加策略
//	@Tags			策略管理
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.AddPolicyReq	true	"添加策略"
//	@Success		200		{string}	string					"{"success":true,"data":{},"msg":"添加成功"}"
//	@Router			/admin/policy/add [post]
func AddPolicy(c *gin.Context) {
	var req request.AddPolicyReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ok, err := service.AddPolicy(&req)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !ok {
		c.JSON(500, gin.H{
			"error": "策略已存在",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"msg":     "添加成功",
	})
}

// RemovePolicy 删除策略
//
//	@Summary		删除策略
//	@Description	删除策略
//	@Tags			策略管理
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.RemovePolicyReq	true	"删除策略"
//	@Success		200		{string}	string					"{"success":true,"data":{},"msg":"删除成功"}"
//	@Router			/admin/policy/remove [post]
func RemovePolicy(c *gin.Context) {
	var req request.RemovePolicyReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ok, err := service.RemovePolicy(&req)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !ok {
		c.JSON(500, gin.H{
			"error": "策略不存在",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"msg":     "删除成功",
	})
}

// UpdatePolicy 修改策略
//
//	@Summary		修改策略
//	@Description	修改策略
//	@Tags			策略管理
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.UpdatePolicyReq	true	"修改策略"
//	@Success		200		{string}	string					"{"success":true,"data":{},"msg":"修改成功"}"
//	@Router			/admin/policy/update [post]
func UpdatePolicy(c *gin.Context) {
	var req request.UpdatePolicyReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ok, err := service.UpdatePolicy(&req)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !ok {
		c.JSON(500, gin.H{
			"error": "策略不存在",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"msg":     "修改成功",
	})
}

// GetPolicyList 获取策略
// summary		获取策略
// Description	获取策略
// Tags			策略管理
// Accept		json
// Produce		json
// Param		request	body		request.GetPolicyReq	true	"获取策略"
// Success		200		{string}	string					"{"success":true,"data":{},"msg":"获取成功"}"
// Router		/admin/policy/list [post]
func GetPolicyList(c *gin.Context) {
	policy := service.GetPolicy()
	c.JSON(200, gin.H{
		"success": true,
		"data":    policy,
		"msg":     "获取成功",
	})
}

// GetPolicyByRole 通过角色获取策略
// summary		通过角色获取策略
// Description	通过角色获取策略
// Tags			策略管理
// Accept		json
// Produce		json
// Param		request	body		request.GetPolicyByRoleReq	true	"通过角色获取策略"
// Success		200		{string}	string					"{"success":true,"data":{},"msg":"获取成功"}"
// Router		/admin/policy/getByRole [post]
func GetPolicyByRole(c *gin.Context) {
	var req request.GetPolicyByRoleReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	policy := service.GetPolicyByRole(&req)
	c.JSON(200, gin.H{
		"success": true,
		"data":    policy,
		"msg":     "获取成功",
	})
}
