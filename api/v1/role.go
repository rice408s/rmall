package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"rmall/model/request"
	"rmall/model/response"
	"rmall/service"
)

// AddRole @Summary AddRole
//
//	@Description	添加角色
//	@Tags			角色
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.AddRoleReq		true	"添加角色"
//	@Success		200		{object}	response.AddRoleResp	"添加成功"
//	@Router			/role/add [post]
func AddRole(c *gin.Context) {
	var req request.AddRoleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := service.AddRole(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.AddRoleResp{Id: id})
}

// UpdateRole @Summary UpdateRole
//
//	@Description	更新角色
//	@Tags			角色
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.UpdateRoleReq	true	"更新角色"
//	@Success		200		{object}	response.UpdateRoleResp	"更新成功"
//	@Router			/role/update [post]
func UpdateRole(c *gin.Context) {
	var req request.UpdateRoleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := service.UpdateRole(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.UpdateRoleResp{Id: req.Id})
}

// DeleteRole
//
//	@Summary		删除角色
//	@Description	删除角色
//	@Tags			角色
//	@Accept			json
//	@Produce		json
//	@Param			request	body	request.DeleteRoleReq	true	"角色id"
//	@Success		200	{object}	response.DeleteRoleResp	"删除成功"
//	@Router			/role/delete [post]
func DeleteRole(c *gin.Context) {
	var req request.DeleteRoleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("err:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println("req:", req)
	err := service.DeleteRole(req.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.DeleteRoleResp{Id: req.Id})
}

// GetRoleList 查询角色列表
//
//	@Summary		查询角色列表
//	@Description	查询角色列表
//	@Tags			角色
//	@Accept			json
//	@Produce		json
//	@Param			request	query		request.GetRoleListReq		true	"查询角色列表"
//	@Success		200		{object}	response.GetRoleListResp	"查询成功"
//	@Router			/role/list [get]
func GetRoleList(c *gin.Context) {
	var req request.GetRoleListReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println("req:", req)

	list, err := service.FindRoleByPage(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.GetRoleListResp{Total: len(list), List: list})
}
