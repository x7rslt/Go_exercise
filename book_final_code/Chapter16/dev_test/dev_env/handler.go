package main

import (
	"dev_env/model"
	"dev_env/repository"
	"dev_env/service"
	"fmt"
)

var (
	DB *gorm.DB
)

func initHandler() {
	fmt.Println("Handler init")
	CommentHandler = handler.CommentHandler{
		Srv: &service.CommentService{
			Repo: &respository.CommentRepo{
				DB: &model.DataBase{
					MyDB: DB,
				},
			},
		},
	}
}
