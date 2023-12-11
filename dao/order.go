package dao

import (
	"encoding/json"
	"errors"
	"fmt"
	"rmall/global"
	"rmall/model"
)

func AddOrder(order *model.Order) (id int, err error) {
	insertStr := "insert into `order` (uid,pid,amount) values(?,?,?)"         // 插入语句
	res, err := global.DB.Exec(insertStr, order.Uid, order.Pid, order.Amount) //执行插入语句
	if err != nil {
		return 0, err
	}
	insertId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	// 更新订单缓存
	key := fmt.Sprintf("orders:id:%d", insertId)
	err = global.RedisCli.Del(ctx, key).Err()
	if err != nil {
		return 0, errors.New("删除订单缓存失败")
	}

	return int(insertId), err //返回插入成功的id

}

// FindOrderById 根据id查找订单
func FindOrderById(id int) (order *model.Order, err error) {
	// 生成一个唯一的键，例如使用"orders:id:" + id
	key := fmt.Sprintf("orders:id:%d", id)
	// 尝试从Redis中获取键对应的数据
	data, err := global.RedisCli.Get(ctx, key).Result()
	if err == nil {
		// 数据存在于Redis中，解析它
		order = &model.Order{}
		err = json.Unmarshal([]byte(data), order)
		if err != nil {
			return nil, err
		}
		fmt.Println("从Redis中获取数据")
		return order, nil
	}
	// 数据不存在于Redis中，从数据库中获取
	queryStr := "select * from `order` where id=?"
	order = &model.Order{}
	err = global.DB.Get(order, queryStr, id)
	if err != nil {
		return nil, err
	}
	fmt.Println("从数据库中获取数据")
	// 将获取的数据存储到Redis中以备后用
	newData, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}
	err = global.RedisCli.Set(ctx, key, newData, 0).Err()
	if err != nil {
		return nil, err
	}
	return order, nil
}

// FindOrderByPage 分页查找订单
func FindOrderByPage(offset, pageSize int) (orders []*model.Order, err error) {
	// 生成一个唯一的键，例如使用"orders:page:" + offset + ":" + pageSize
	key := fmt.Sprintf("orders:page:%d:%d", offset, pageSize)

	// 尝试从Redis中获取键对应的数据
	data, err := global.RedisCli.Get(ctx, key).Result()
	if err == nil {
		// 数据存在于Redis中，解析它
		err = json.Unmarshal([]byte(data), &orders)
		if err != nil {
			return nil, err
		}
		fmt.Println("从Redis中获取数据")
		return orders, nil
	}
	// 数据不存在于Redis中，从数据库中获取
	queryStr := "select * from `order` limit ? offset ?"
	orders = []*model.Order{}
	err = global.DB.Select(&orders, queryStr, pageSize, offset)
	if err != nil {
		return nil, err
	}
	fmt.Println("从数据库中获取数据")
	// 将获取的数据存储到Redis中以备后用
	newData, err := json.Marshal(orders)
	if err != nil {
		return nil, err
	}
	err = global.RedisCli.Set(ctx, key, newData, 0).Err()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// UpdateOrder 更新订单
func UpdateOrder(order *model.Order) (err error) {
	updateStr := "update `order` set uid=?,pid=?,amount=? where id=?"
	_, err = global.DB.Exec(updateStr, order.Uid, order.Pid, order.Amount, order.Id)
	if err != nil {
		return err
	}
	// 更新订单缓存
	key := fmt.Sprintf("orders:id:%d", order.Id)
	err = global.RedisCli.Del(ctx, key).Err()
	if err != nil {
		return errors.New("删除订单缓存失败")
	}
	return nil
}

// DeleteOrder 删除订单
func DeleteOrder(id int) (err error) {
	// 开启数据库事务
	tx, err := global.DB.Begin()
	if err != nil {
		return err
	}

	// 查询订单是否存在
	order, err := FindOrderById(id)
	if err != nil {
		return tx.Rollback()
	}

	// 恢复商品库存
	product, err := FindProductById(int(order.Pid))
	if err != nil {
		return tx.Rollback()

	}
	product.Stock += order.Amount
	err = UpdateProduct(product)
	if err != nil {
		return tx.Rollback()
	}

	// 删除订单
	deleteStr := "DELETE FROM `order` WHERE id=?"
	_, err = tx.Exec(deleteStr, id)
	if err != nil {
		return tx.Rollback()
	}

	// 更新订单缓存
	key := fmt.Sprintf("orders:id:%d", id)
	err = global.RedisCli.Del(ctx, key).Err()
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return errors.New("回滚失败")
		}
		return errors.New("删除订单缓存失败")
	}

	// 提交事务
	err = tx.Commit()
	if err != nil {
		return tx.Rollback()

	}

	return nil
}

// FindOrderByUid 根据用户id查找订单
func FindOrderByUid(uid int) (orders []*model.Order, err error) {
	queryStr := "select * from `order` where uid=?"
	orders = []*model.Order{}
	err = global.DB.Select(&orders, queryStr, uid)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// FindOrderByPid 根据商品id查找订单
func FindOrderByPid(pid int) (orders []*model.Order, err error) {
	queryStr := "select * from `order` where pid=?"
	orders = []*model.Order{}
	err = global.DB.Select(&orders, queryStr, pid)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
