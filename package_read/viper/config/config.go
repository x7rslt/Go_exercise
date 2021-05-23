package config

import (
	"github.com/fsnotify/fsnotify"
	//"github.com/i-coder-robot/book_final_code/Chapter16/MyLog"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Name string
}

func Init(name string) error {
	c := Config{
		Name: name,
	}
	//初始化配置文件
	if err := c.initConfig(); err != nil {
		log.Println(err)
		return err
	}
	c.watchConfig()
	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name) //如果指定配置文件，解析
	} else {
		viper.AddConfigPath("./conf/")
		viper.SetConfigName("app_config")
	}
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

//监控配置文件变化
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:%s", e.Name)
	})
}
