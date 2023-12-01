package request

type CreateProductReq struct {
	Name  string `json:"name" binding:"required"`
	Price int    `json:"price" binding:"required"`
	Stock int    `json:"stock" binding:"required"`
	// 产品状态 0-下架 1-上架
	Status int    `json:"status" binding:"required"`
	Desc   string `json:"desc" binding:"required"`
}

// UpdateProductReq 更新产品请求
type UpdateProductReq struct {
	Id    int    `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Price int    `json:"price" binding:"required"`
	Stock int    `json:"stock" binding:"required"`
	// 产品状态 0-下架 1-上架
	Status int    `json:"status" binding:"required"`
	Desc   string `json:"desc" binding:"required"`
}

// GetProductListReq 获取产品列表请求
type GetProductListReq struct {
	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}

// GetProductByIdReq 根据id获取产品请求
type GetProductByIdReq struct {
	Id int `json:"id" form:"id"`
}

// GetProductByNameReq 根据产品名称获取产品
type GetProductByNameReq struct {
	Name string `json:"name" form:"name"`
}

type DeleteProductReq struct {
	Id int `json:"id" form:"id"`
}
