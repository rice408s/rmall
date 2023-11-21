package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"rmall/global"
)

func LoadConfig(path string) {
	viper.SetConfigFile(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	//添加配置文件所在的路径，注意在Linux环境下%GOPATH要替换为$GOPATH
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("viper.ReadInConfig failed,err:", err)
		return
	}
	// 参数使用指针
	err := viper.Unmarshal(&global.Config)
	if err != nil {
		fmt.Println("viper.Unmarshal failed,err:", err)
		return
	}
}
