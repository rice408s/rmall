package main

import (
	_ "rmall/docs"
	"rmall/initialize"
)

// 添加注释以描述 server 信息
//	@title			商城系统
//	@version		1.0
//	@description	商城系统接口文档

//	@contact.name	rice408
//	@contact.url	http://47.108.226.33/
//	@contact.email	hongwei.wang408@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

func main() {
	initialize.Run()
}
