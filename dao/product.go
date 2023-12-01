package dao

import (
	"fmt"
	"rmall/global"
	"rmall/model"
)

// AddProduct 添加商品
func AddProduct(product *model.Product) (id int, err error) {
	insertStr := "insert into product(name, `desc`, stock, price, status,create_time,update_time) values(?,?,?,?,?,?,?)"
	fmt.Println(insertStr)
	fmt.Println(product)
	result, err := global.DB.Exec(insertStr, product.Name, product.Desc, product.Stock, product.Price, product.Status, product.CreateTime, product.UpdateTime)
	if err != nil {
		return 0, err
	}
	insertId, err := result.LastInsertId()
	return int(insertId), err
}

// FindProductById 根据id查找商品
func FindProductById(id int) (product *model.Product, err error) {
	queryStr := "select * from product where id=?"
	product = &model.Product{}
	err = global.DB.Get(product, queryStr, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// FindProductByPage 分页查找商品
func FindProductByPage(offset, pageSize int) (products []*model.Product, err error) {
	queryStr := "select * from product limit ? offset ?"
	products = []*model.Product{}
	err = global.DB.Select(&products, queryStr, pageSize, offset)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// FindProductByName 根据名称查找商品列表,模糊查询
func FindProductByName(name string) (products []*model.Product, err error) {
	queryStr := "select * from product where name like ?"
	products = []*model.Product{}
	err = global.DB.Select(&products, queryStr, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	return products, nil
}

// UpdateProduct 更新商品
func UpdateProduct(product *model.Product) (err error) {
	updateStr := "update product set name=?, `desc`=?, stock=?, price=?, status=? where id=?"
	_, err = global.DB.Exec(updateStr, product.Name, product.Desc, product.Stock, product.Price, product.Status, product.Id)
	return err
}

// DeleteProductById 根据id删除商品
func DeleteProductById(id int) (err error) {
	deleteStr := "delete from product where id=?"
	_, err = global.DB.Exec(deleteStr, id)
	return err
}
