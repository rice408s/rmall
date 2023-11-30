package global

import (
	"github.com/casbin/casbin/v2"
	"github.com/jmoiron/sqlx"
	"rmall/config"
)

// 定义全局变量
var (
	Config config.Config
	Db     *sqlx.DB
	E      *casbin.Enforcer
)
