package model

import (
	"time"
)

// Product 定义产品结构体
type Product struct {
	Id    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Desc  string `json:"desc" db:"desc"`   // 产品描述
	Stock int    `json:"stock" db:"stock"` // 库存
	Price int    `json:"price" db:"price"` // 价格
	// 产品状态 0-下架 1-上架
	Status     int       `json:"status" db:"status"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
	UpdateTime time.Time `json:"update_time" db:"update_time"`
}
