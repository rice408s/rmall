package initialize

import (
	"rmall/router"
)

func Run() {
	loadConfig()
	initMysql()
	initRedis()
	initCasbin()
	initAliPay()
	timedTask()
	router.Route()
}
