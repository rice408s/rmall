package initialize

import "rmall/router"

func Run() {
	LoadConfig("./config.yaml")
	Mysql()
	router.Route()
}
