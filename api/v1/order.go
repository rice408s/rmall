package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"rmall/model/request"
	"rmall/model/response"
	"rmall/service"
)

// CreateOrder 创建订单
//
//	@Summary		创建订单
//	@Description	创建订单
//	@Tags			订单
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.CreateOrderReq true	"创建订单"
//	@Success		200		{object}	response.CreateOrderResp ""
//	@Router			/order/create [post]
func CreateOrder(c *gin.Context) {
	var req request.CreateOrderReq
	claims, _ := c.Get("user")
	userClaims := claims.(jwt.MapClaims)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	req.Uid = int64(userClaims["id"].(float64))
	//创建订单
	id, err := service.CreateOrder(&req)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(http.StatusOK, response.CreateOrderResp{Id: id})

}

// UpdateOrder 修改订单
// @Summary 修改订单
// @Description 修改订单
// @Tags 订单
// @Accept json
// @Produce json
// @Param request body request.UpdateOrderReq true "修改订单"
// @Success 200 {object} response.UpdateOrderResp ""
// @Router /user/order/update [post]
func UpdateOrder(c *gin.Context) {
	var req request.UpdateOrderReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	err = service.UpdateOrder(&req)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(http.StatusOK, response.UpdateOrderResp{
		Id: int(req.Id),
	})
}

// GetOrderById 根据id获取订单
// @Summary 根据id获取订单
// @Description 根据id获取订单
// @Tags 订单
// @Accept json
// @Produce json
// @Param request body request.GetOrderByIdReq true "根据id获取订单"
// @Success 200 {object} response.GetOrderByIdResp
// @Router /user/order/get [post]
func GetOrderById(c *gin.Context) {
	var req request.GetOrderByIdReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	order, err := service.FindOrderById(req.Id)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(http.StatusOK, response.GetOrderByIdResp{
		Order: order,
	})
}

// GetOrderList 分页获取订单
// @Summary 获取订单列表
// @Description 分页获取订单
// @Tags 订单
// @Accept json
// @Produce json
// @Param request body request.GetOrderListReq true "分页获取订单"
// @Success 200 {object} response.GetOrderListResp
// @Router /user/order/list [post]
func GetOrderList(c *gin.Context) {
	var req request.GetOrderListReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	orders, err := service.FindOrderByPage(&req)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(http.StatusOK, response.GetOrderListResp{
		Total: len(orders),
		List:  orders,
	})
}

// DeleteOrder 删除订单
// @Summary 删除订单
// @Description 删除订单
// @Tags 订单
// @Accept json
// @Produce json
// @Param request body request.DeleteOrderReq true "删除订单"
// @Success 200 {object} response.DeleteOrderResp ""
// @Router /user/order/delete [post]
func DeleteOrder(c *gin.Context) {
	var req request.DeleteOrderReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	err = service.DeleteOrder(req.Id)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(http.StatusOK, response.DeleteOrderResp{
		Id: req.Id,
	})
}

// GetOrderByUid 通过uid获取订单
// @Summary 通过用户获取订单
// @Description 通过uid获取订单
// @Tags 订单
// @Accept json
// @Produce json
// @Param request body request.GetOrderByUidReq true "通过uid获取订单"
// @Success 200 {object} response.GetOrderByUidResp
// @Router /user/order/getByUid [post]
// @Router /admin/order/getByUid [post]
func GetOrderByUid(c *gin.Context) {
	var req request.GetOrderByUidReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(399, err.Error())
		return
	}

	claims, _ := c.Get("user")
	userClaims := claims.(jwt.MapClaims)
	req.Uid = int(userClaims["id"].(float64))

	orders, err := service.FindOrderByUid(req.Uid)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(http.StatusOK, response.GetOrderByUidResp{
		Total: len(orders),
		List:  orders,
	})
}

// GetOrderByPid 通过商品获取订单
// @Summary 通过商品获取订单
// @Description 通过pid获取订单
// @Tags 订单
// @Accept json
// @Produce json
// @Param request body request.GetOrderByPidReq true "通过pid获取订单"
// @Success 200 {object} response.GetOrderByPidResp
// @Router /user/order/getByPid [post]
func GetOrderByPid(c *gin.Context) {
	var req request.GetOrderByPidReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(399, err.Error())
		return
	}
	orders, err := service.FindOrderByPid(req.Pid)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(http.StatusOK, response.GetOrderByPidResp{
		Total: len(orders),
		List:  orders,
	})
}
