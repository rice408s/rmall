package initialize

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"rmall/global"
)

func initMysql() {
	//拼接dsn
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		global.Config.Mysql.Username,
		global.Config.Mysql.Password,
		global.Config.Mysql.Host,
		global.Config.Mysql.Port,
		global.Config.Mysql.Dbname,
		global.Config.Mysql.Params,
	)
	fmt.Println(dsn)
	var err error
	global.DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("sqlx.Connect failed,err:", err)
		return
	}
	// 最大空闲连接数
	global.DB.SetMaxIdleConns(20)
	// 最大打开连接数
	global.DB.SetMaxOpenConns(20)
}

func initRedis() {
	global.RedisCli = redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.Addr,
		Password: global.Config.Redis.Password,
		DB:       global.Config.Redis.DB,
	})
	if global.RedisCli == nil {
		fmt.Println("redis.NewClient failed")
		return
	}
	ctx := context.Background()
	get := global.RedisCli.Get(ctx, "ss")
	fmt.Println(get.Val())
}
