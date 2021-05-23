package main

import (
	"confset/conf"
	"confset/config"
	"fmt"
	"github.com/spf13/viper"
)

func initViper() {
	if err := config.Init(""); err != nil {
		panic(err)
	}
}

func init() {
	initViper()
}

func main() {

	conf := &conf.DBConf{
		Host:     viper.GetString("database.host"),
		User:     viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		DbName:   viper.GetString("database.name"),
	}
	fmt.Println(conf)
}
