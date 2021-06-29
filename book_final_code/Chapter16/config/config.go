package config

import (
	"food/MyLog"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Name string
}

func Init(name string)error{
	c := Config{
		Name:name,
	}
	if err := c.initConfig();err != nil{
		MyLog.Log.Error(err)
		return err
	}
	c.watchConfig()
	return nil
}

func (c *Config)initConfig()error{
	if c.Name != ""{
		viper.SetConfigFile(c.Name)
	}else{
		viper.AddConfigPath("./conf/")
		viper.SetConfigName("app_config")
	}
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig();err !=nil{
		return err

	}
	return nil
}

func (c *Config)watchConfig(){
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event){
		log.Printf("Config file changed:#{e.Name}")
	})
}


