package api

import (
	"github.com/gin-gonic/gin"
	"rmall/model/request"
)

// CreateOrder 创建订单
//
//	@Summary		创建订单
//	@Description	创建订单创建订单
//	@Tags			订单
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.CreateOrderReq
//	@Success		200		{object}	response.CreateOrderResp
//	@Router			/order/create [post]
func CreateOrder(c *gin.Context) {
	var req request.CreateOrderReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	//创建订单

}
