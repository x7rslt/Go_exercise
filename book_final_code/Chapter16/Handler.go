package main

import (
	"fmt"
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
		Password: viper.GetString("databse.password"),
		DbName:   viper.GetString("databse.name"),
	}
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
