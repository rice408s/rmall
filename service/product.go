package service

import (
	"rmall/dao"
	"rmall/model"
	"rmall/model/request"
	"time"
)

// CreateProduct 创建商品
func CreateProduct(req *request.CreateProductReq) (int, error) {
	// 参数校验
	if req.Name == "" || req.Desc == "" || req.Stock == 0 || req.Price == 0 {
		return 0, nil
	}
	// 保存商品信息
	product := &model.Product{
		Name:       req.Name,
		Desc:       req.Desc,
		Stock:      req.Stock,
		Price:      req.Price,
		Status:     req.Status,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	//插入数据库
	return dao.AddProduct(product)

}

// FindProductById 根据id查找商品
func FindProductById(id int) (product *model.Product, err error) {
	return dao.FindProductById(id)
}

// FindProductByPage 分页查找商品
func FindProductByPage(req *request.GetProductListReq) (products []*model.Product, err error) {
	start := (req.PageNum - 1) * req.PageSize
	return dao.FindProductByPage(start, req.PageSize)
}

// FindProductByName 根据名称查找商品,模糊查询
func FindProductByName(name string) (product []*model.Product, err error) {
	return dao.FindProductByName(name)
}

// UpdateProduct 更新商品
func UpdateProduct(req *request.UpdateProductReq) (err error) {
	// 参数校验
	if req.Name == "" || req.Desc == "" || req.Stock == 0 || req.Price == 0 {
		return nil
	}
	//查询商品是否存在
	_, err = dao.FindProductById(req.Id)
	if err != nil {
		return err
	}

	product := &model.Product{
		Id:         req.Id,
		Name:       req.Name,
		Desc:       req.Desc,
		Stock:      req.Stock,
		Price:      req.Price,
		UpdateTime: time.Now(),
	}
	return dao.UpdateProduct(product)
}

// DeleteProduct 删除商品
func DeleteProduct(req *request.DeleteProductReq) (err error) {

	_, err = dao.FindProductById(req.Id)
	if err != nil {
		return err
	}

	return dao.DeleteProductById(req.Id)
}
