package service

import (
	"errors"
	"rmall/dao"
	"rmall/global"
	"rmall/model"
	"rmall/model/request"
	"time"
)

// CreateOrder 创建订单
func CreateOrder(req *request.CreateOrderReq) (int, error) {
	// 参数校验
	if req.Uid == 0 || req.Pid == 0 || req.Amount == 0 {
		return 0, errors.New("参数错误")
	}
	//查询商品是否存在
	product, err := dao.FindProductById(int(req.Pid))
	if err != nil {
		return 0, errors.New("商品不存在")
	}
	//计算商品库存
	if product.Stock < req.Amount {
		return 0, errors.New("库存不足")
	}

	//查询用户是否存在
	_, err = dao.FindUserById(int(req.Uid))
	if err != nil {
		return 0, errors.New("用户不存在")
	}
	//添加事务
	var tx, _ = global.DB.Beginx()

	// 保存订单信息
	order := &model.Order{
		Uid:        req.Uid,
		Pid:        req.Pid,
		Amount:     req.Amount,
		Status:     0,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	//修改商品库存
	product.Stock = product.Stock - req.Amount
	err = dao.UpdateProduct(product)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return 0, errors.New("回滚失败")
		}
		return 0, errors.New("修改商品库存失败")
	}
	//插入订单
	id, err := dao.AddOrder(order)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return 0, errors.New("回滚失败")
		}
		return 0, errors.New("插入订单失败")
	}
	//提交事务
	err = tx.Commit()
	if err != nil {
		return 0, errors.New("提交事务失败")
	}
	return id, nil
}

// FindOrderById 根据id查找订单
func FindOrderById(id int) (order *model.Order, err error) {
	return dao.FindOrderById(id)
}

// FindOrderByPage 分页查找订单
func FindOrderByPage(req *request.GetOrderListReq) (orders []*model.Order, err error) {
	start := (req.PageNum - 1) * req.PageSize
	return dao.FindOrderByPage(start, req.PageSize)
}

// UpdateOrder 更新订单
func UpdateOrder(req *request.UpdateOrderReq) (err error) {
	// 参数校验
	if req.Uid == 0 || req.Pid == 0 || req.Amount == 0 {
		return nil
	}
	//查询订单是否存在
	_, err = dao.FindOrderById(int(req.Id))
	if err != nil {
		return err
	}

	order := &model.Order{
		Id:         req.Id,
		Uid:        req.Uid,
		Pid:        req.Pid,
		Amount:     req.Amount,
		UpdateTime: time.Now(),
	}
	return dao.UpdateOrder(order)
}

// DeleteOrder 删除订单
func DeleteOrder(id int) (err error) {
	//查询订单是否存在
	_, err = dao.FindOrderById(int(id))
	if err != nil {
		return err
	}
	return dao.DeleteOrder(id)
}

// FindOrderByUid 根据用户id查找订单
func FindOrderByUid(uid int) (orders []*model.Order, err error) {
	return dao.FindOrderByUid(uid)
}

// FindOrderByPid 根据商品id查找订单
func FindOrderByPid(pid int) (orders []*model.Order, err error) {
	return dao.FindOrderByPid(pid)
}
