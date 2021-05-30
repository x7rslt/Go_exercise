package repository

import (
	"dev_env/model"
	"fmt"
	"gorm.io/gorm"
	"log"
	//"net/http"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/schema"
)

type CommentRepo struct {
	DB model.DataBase
}

func search() *gorm.DB {
	dsn := "root:Xss8271329.X@tcp(167.99.155.35:3306)/food_app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
	}})
	fmt.Println(db)
	if err != nil {
		fmt.Println(err)

	}
	return db
}
func (c *CommentRepo) GetCommentList() []model.Comment {
	fmt.Println("Repository start.")

	var comment1 model.Comment
	fmt.Println("Find 1 comment:", comment1)
	db1 := search()
	db1.First(&comment1)

	fmt.Println("Find 1 comment:", comment1)
	var commentList []model.Comment
	c.DB.MyDB.Find(&commentList)
	log.Println("Repositroy result:", commentList)
	fmt.Println("Repository no probelom")
	return commentList
}
