package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type MySQL struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	dns  string
}

func InitDB() (*sqlx.DB, error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/rmall?charset=utf8mb4&parseTime=True"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("connect DB failed, err:", err)
		return nil, err
	}

	// 最大空闲连接数
	db.SetMaxIdleConns(20)
	// 最大打开连接数
	db.SetMaxOpenConns(20)
	return db, nil
}
