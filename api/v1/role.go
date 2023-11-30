package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"rmall/global"
	"rmall/model/request"
	"rmall/model/response"
	"rmall/service"
	"strconv"
)

// AddRole
//
//	@Summary		添加角色
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
//	@summary		更新角色
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
	// 更新角色的同时不用更新策略，因为策略使用的是角色id，角色id不会变

	c.JSON(http.StatusOK, response.UpdateRoleResp{Id: req.Id})
}

// DeleteRole
//
//	@Summary		删除角色
//	@Description	删除角色
//	@Tags			角色
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.DeleteRoleReq	true	"角色id"
//	@Success		200		{object}	response.DeleteRoleResp	"删除成功"
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
	// 删除角色的同时删除角色策略
	ok, err := global.E.RemoveFilteredPolicy(0, strconv.Itoa(req.Id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "删除策略失败",
		})
		return
	}

	// 删除角色的同时删除管理员与角色的关联
	ok, err = service.RemovePolicyByRole(strconv.Itoa(req.Id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "删除管理员关联角色策略失败",
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

// GetRoleListByAdmin
// @Summary 通过管理员id获取角色列表
// @Description 通过管理员id获取角色列表
// @Tags 角色
// @Accept json
// @Produce json
// @Param request query request.GetRoleListByAdminIdReq true "通过管理员id获取角色列表"
// @Success 200 {object} string
// @Router /role/listByAdminId [get]
func GetRoleListByAdmin(c *gin.Context) {
	var req request.GetRoleListByAdminIdReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println("req:", req)

	list, err := service.FindRoleByAdminId(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.GetRoleListByAdminIdResp{Total: len(list), List: list})
}
