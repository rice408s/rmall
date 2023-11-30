package initialize

import "rmall/router"

func Run() {
	loadConfig("./config.yaml")
	mysql()
	initCasbin()
	router.Route()
}
