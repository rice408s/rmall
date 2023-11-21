package initialize

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"rmall/global"
)

func Mysql() {
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
	global.Db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("sqlx.Connect failed,err:", err)
		return
	}
	// 最大空闲连接数
	global.Db.SetMaxIdleConns(20)
	// 最大打开连接数
	global.Db.SetMaxOpenConns(20)
}
