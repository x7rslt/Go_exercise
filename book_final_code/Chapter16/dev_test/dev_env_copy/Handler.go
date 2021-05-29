package main

import (
	"dev_env/handler"
	"dev_env/model"
	"dev_env/repository"
	"dev_env/service"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DB *gorm.DB
)

func initDB() {
	fmt.Println("DB init")
	var err error
	dsn := "root:Xss8271329.X@tcp(167.99.155.35:3306)/food_app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
	}})
	if err != nil {
		fmt.Println(err)
	}
	var comment []model.Comment

	db.Find(&comment)
	fmt.Println(comment)
	fmt.Println("数据库 init 结束...")
}

func initHandler() {
	fmt.Println("Handler init")
	CommentHandler = handler.CommentHandler{
		Srv: &service.CommentService{
			Repo: &repository.CommentRepo{
				DB: model.DataBase{
					MyDB: DB,
				},
			},
		},
	}
	fmt.Println("inithandler :", CommentHandler)
}
