package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rmall/model/request"
	"rmall/model/response"
	"rmall/service"
)

// AddProduct 添加产品
//	@Summary		添加产品
//	@Description	添加产品
//	@Tags			产品
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.CreateProductReq	true	"添加产品"
//	@Success		200		{object}	response.CreateProductResp	"添加成功"
//	@Router			/product/add [post]
func AddProduct(c *gin.Context) {
	var req request.CreateProductReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := service.CreateProduct(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.CreateProductResp{Id: id})
}

// UpdateProduct 更新产品
//	@Summary		更新产品
//	@Description	更新产品
//	@Tags			产品
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.UpdateProductReq	true	"更新产品"
//	@Success		200		{object}	response.UpdateProductResp	"更新成功"
//	@Router			/product/update [post]
func UpdateProduct(c *gin.Context) {
	var req request.UpdateProductReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = service.UpdateProduct(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.UpdateProductResp{Id: req.Id})
}

// GetProductList 获取产品列表
//	@Summary		获取产品列表
//	@Description	获取产品列表
//	@Tags			产品
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.GetProductListReq	true	"获取产品列表"
//	@Success		200		{object}	response.GetProductListResp	"获取成功"
//	@Router			/product/list [post]
func GetProductList(c *gin.Context) {
	var req request.GetProductListReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	products, err := service.FindProductByPage(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.GetProductListResp{Total: len(products), List: products})
}

// DeleteProduct 删除产品
//	@Summary		删除产品
//	@Description	删除产品
//	@Tags			产品
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.DeleteProductReq	true	"删除产品"
//	@Success		200		{object}	response.DeleteProductResp	"删除成功"
//	@Router			/product/delete [post]
func DeleteProduct(c *gin.Context) {
	var req request.DeleteProductReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = service.DeleteProduct(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.DeleteProductResp{Id: req.Id})
}

// GetProductById 根据id获取产品
//	@Summary		根据id获取产品
//	@Description	根据id获取产品
//	@Tags			产品
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.GetProductByIdReq	true	"根据id获取产品"
//	@Success		200		{object}	response.GetProductByIdResp	"获取成功"
//	@Router			/product/get [post]
func GetProductById(c *gin.Context) {
	var req request.GetProductByIdReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	product, err := service.FindProductById(req.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.GetProductByIdResp{Product: product})
}

// GetProductByName 模糊查询
// GetProductByName 根据名称获取产品列表
//	@Summary		根据名称获取产品列表
//	@Description	根据名称获取产品列表
//	@Tags			产品
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.GetProductByNameReq		true	"根据名称获取产品列表"
//	@Success		200		{object}	response.GetProductByNameResp	"获取成功"
//	@Router			/product/getByName [post]
func GetProductByName(c *gin.Context) {
	var req request.GetProductByNameReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	products, err := service.FindProductByName(req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.GetProductByNameResp{Total: len(products), List: products})
}
