package config

import (
	"fmt"
	conf "github.com/Unknwon/goconfig"
)

var appConf *conf.ConfigFile

func init() {
	var err error
	appConf, err = conf.LoadConfigFile("app.conf")
	if err != nil {
		fmt.Println("载入配置文件出错:")
		panic(err)
	}
}

func GetPort() string {
	port, err := appConf.GetValue("app", "port")
	if err != nil {
		fmt.Println("获取配置PORT出错了,使用默认PORT:3000")
		port = "3000"
	}
	return port
}
