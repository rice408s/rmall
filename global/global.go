package global

import (
	"github.com/jmoiron/sqlx"
	"rmall/config"
)

// 定义全局变量
var (
	Config config.Config
	Db     *sqlx.DB
)
