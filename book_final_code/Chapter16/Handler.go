package main

import (
	"fmt"
	"food/conf"
	"food/config"
	"food/handler"
	"food/model"
	"food/repository"
	"food/service/dp_service"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var (
	DB *gorm.DB
)

func initViper() {
	if err := config.Init(""); err != nil {
		panic(err)
	}
}

func initDB() {
	fmt.Println("数据库init")
	var err error
	conf := &conf.DBConf{
		Host:     viper.GetString("database.host"),
		User:     viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		DbName:   viper.GetString("database.name"),
	}
	log.Println(conf)
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8&parseTime=%t&loc=%s",
		conf.User,
		conf.Password,
		conf.Host,
		conf.DbName,
		true,
		"Local")
	log.Println(config)
	DB, err = gorm.Open("mysql", config)
	if err != nil {
		log.Fatalf("connect error:%v\n", err)
	}
	DB.SingularTable(true) //待重新配置
	fmt.Println("数据库init完成...")

}
func initHandler() {
	HotelDetailHandler = handler.HotelDetailHandler{
		Srv: &dp_service.HotelService{
			Repo: &repository.HotelRepo{
				DB: model.DataBase{
					MyDB: DB,
				},
			},
			TeamRepo: &repository.TeamRepo{
				DB: model.DataBase{
					MyDB: DB,
				},
			},
			SuggestFoodRepo: &repository.SuggestFoodRepo{
				DB: model.DataBase{
					MyDB: DB,
				},
			},
			CommentTagRepo: &repository.CommentTagRepo{
				DB: model.DataBase{
					MyDB: DB,
				},
			},
			CommentRepo: &repository.CommentRepo{
				DB: model.DataBase{
					MyDB: DB,
				},
			},
			MarketRepo: &repository.MarketRepo{
				DB: model.DataBase{
					MyDB: DB,
				},
			},
		},
	}
}
