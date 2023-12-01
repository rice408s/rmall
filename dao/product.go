package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"rmall/global"
	"rmall/model"
	"time"
)

var ctx = context.Background()

// AddProduct 添加商品
func AddProduct(product *model.Product) (id int, err error) {
	insertStr := "insert into product(name, `desc`, stock, price, status,create_time,update_time) values(?,?,?,?,?,?,?)"
	fmt.Println(product)
	result, err := global.DB.Exec(insertStr, product.Name, product.Desc, product.Stock, product.Price, product.Status, product.CreateTime, product.UpdateTime)
	if err != nil {
		return 0, err
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	// 更新缓存
	key := fmt.Sprintf("products:id:%d", insertId)
	err = global.RedisCli.Del(ctx, key).Err()
	if err != nil {
		return 0, err
	}

	return int(insertId), UpdateProductCache()
}

// FindProductById 根据id查找商品
func FindProductById(id int) (product *model.Product, err error) {
	// 生成一个唯一的键，例如使用"products:id:" + id
	key := fmt.Sprintf("products:id:%d", id)
	// 尝试从Redis中获取键对应的数据
	data, err := global.RedisCli.Get(ctx, key).Result()
	if err == nil {
		// 数据存在于Redis中，解析它
		product = &model.Product{}
		err = json.Unmarshal([]byte(data), product)
		if err != nil {
			return nil, err
		}
		fmt.Println("从Redis中获取数据")
		return product, nil
	}
	// 数据不存在于Redis中，从数据库中获取
	queryStr := "select * from product where id=?"
	product = &model.Product{}
	err = global.DB.Get(product, queryStr, id)
	if err != nil {
		return nil, err
	}
	fmt.Println("从数据库中获取数据")
	// 将获取的数据存储到Redis中以备后用
	newData, err := json.Marshal(product)
	if err != nil {
		return nil, err
	}
	err = global.RedisCli.Set(ctx, key, newData, time.Hour).Err() // 设置缓存一个小时后过期
	return product, nil
}

// FindProductByPage 分页查找商品
func FindProductByPage(offset, pageSize int) (products []*model.Product, err error) {
	// 生成一个唯一的键，例如使用"products:page:" + offset + ":" + pageSize
	key := fmt.Sprintf("products:page:%d:%d", offset, pageSize)

	// 尝试从Redis中获取键对应的数据
	data, err := global.RedisCli.Get(ctx, key).Result()
	if err == nil {
		// 数据存在于Redis中，解析它
		err = json.Unmarshal([]byte(data), &products)
		if err != nil {
			return nil, err
		}
		fmt.Println("从Redis中获取数据")
		return products, nil
	}

	// 数据不存在于Redis中，从数据库中获取
	queryStr := "select * from product limit ? offset ?"
	products = []*model.Product{}
	err = global.DB.Select(&products, queryStr, pageSize, offset)
	if err != nil {
		return nil, err
	}
	fmt.Println("从数据库中获取数据")
	// 将获取的数据存储到Redis中以备后用
	newData, err := json.Marshal(products)

	if err != nil {
		return nil, err
	}
	err = global.RedisCli.Set(ctx, key, newData, time.Hour).Err() // 设置缓存一个小时后过期
	if err != nil {
		return nil, err
	}

	return products, nil
}

// FindProductByName 根据名称查找商品列表,模糊查询
func FindProductByName(name string) (products []*model.Product, err error) {
	//生成一个唯一的键，例如使用"products:name:" + name
	key := fmt.Sprintf("products:name:%s", name)
	// 尝试从Redis中获取键对应的数据
	data, err := global.RedisCli.Get(ctx, key).Result()
	if err == nil {
		// 数据存在于Redis中，解析它
		err = json.Unmarshal([]byte(data), &products)
		if err != nil {
			return nil, err
		}
		fmt.Println("从Redis中获取数据")
		return products, nil
	}
	// 数据不存在于Redis中，从数据库中获取
	queryStr := "select * from product where name like ?"
	products = []*model.Product{}
	err = global.DB.Select(&products, queryStr, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	fmt.Println("从数据库中获取数据")
	// 将获取的数据存储到Redis中以备后用
	newData, err := json.Marshal(products)
	if err != nil {
		return nil, err
	}
	err = global.RedisCli.Set(ctx, key, newData, time.Hour).Err() // 设置缓存一个小时后过期
	return products, nil
}

// UpdateProduct 更新商品
func UpdateProduct(product *model.Product) (err error) {
	updateStr := "update product set name=?, `desc`=?, stock=?, price=?, status=? where id=?"
	_, err = global.DB.Exec(updateStr, product.Name, product.Desc, product.Stock, product.Price, product.Status, product.Id)
	if err != nil {
		return err
	}
	// 更新缓存
	key := fmt.Sprintf("products:id:%d", product.Id)
	err = global.RedisCli.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	return UpdateProductCache()
}

// DeleteProductById 根据id删除商品
func DeleteProductById(id int) (err error) {
	deleteStr := "delete from product where id=?"
	_, err = global.DB.Exec(deleteStr, id)
	if err != nil {
		return err
	}
	// 更新缓存
	key := fmt.Sprintf("products:id:%d", id)
	err = global.RedisCli.Del(ctx, key).Err()

	if err != nil {
		return err
	}
	fmt.Println("删除缓存")
	return UpdateProductCache()
}

// UpdateProductCache 更新分页查询缓存和模糊查询缓存
func UpdateProductCache() (err error) {
	//更新分页缓存

	//先查询redis中是否有匹配的数据
	keys, err := global.RedisCli.Keys(ctx, "products:page:*").Result()
	if err != nil {
		return err
	}
	if len(keys) > 0 {
		global.RedisCli.Del(ctx, keys...)
	}

	//更新模糊查询缓存
	keys, err = global.RedisCli.Keys(ctx, "products:name:*").Result()

	if err != nil {
		return err
	}
	if len(keys) > 0 {
		del := global.RedisCli.Del(ctx, keys...)
		if del.Err() != nil {
			return del.Err()
		}
	}
	return nil
}
