package response

import "rmall/model"

// CreateProductResp 创建商品
type CreateProductResp struct {
	Id int `json:"id"`
}

// UpdateProductResp 更新商品
type UpdateProductResp struct {
	Id int `json:"id"`
}

// GetProductListResp 获取商品列表
type GetProductListResp struct {
	Total int              `json:"total"`
	List  []*model.Product `json:"list"`
}

// GetProductByIdResp 根据id获取商品
type GetProductByIdResp struct {
	*model.Product
}

// GetProductByNameResp 根据名称获取商品列表
type GetProductByNameResp struct {
	Total int              `json:"total"` // 总数
	List  []*model.Product `json:"list"`  // 产品列表
}

type DeleteProductResp struct {
	Id int `json:"id"`
}
