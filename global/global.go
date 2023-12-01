package global

import (
	"github.com/casbin/casbin/v2"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"rmall/config"
)

// 定义全局变量
var (
	Config config.Config
	DB     *sqlx.DB
	E      *casbin.Enforcer

	RedisCli *redis.Client
)
