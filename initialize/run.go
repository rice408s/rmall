package initialize

import (
	"rmall/router"
)

func Run() {
	loadConfig("./config.yaml")
	mysql()
	initRedis()
	initCasbin()
	timedTask()
	router.Route()
}
