package initialize

func Run() {
	LoadConfig("./config.yaml")
	Mysql()
}
