package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	//"net/http"
	"gorm.io/gorm/schema"
)

type Comment struct {
	//评论ID
	Commentid string `json:"commentid"`
	//评论
	Comment string `json:"comment"`
	//评论人
	Accountname string `json:"accountname"`
	//评论人头像
	Accountpic string `json:"accountpic"`
	//评分
	Star int `json:"star"`
	//人均消费
	Averageperson int `json:"averageperson"`
	//是否是优质点评
	Isgood int `json:"isgood"`
}

func main() {

	dsn := "root:mysqltest110@tcp(114.116.230.93:3306)/food_app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
	}})
	fmt.Println(db)
	if err != nil {
		fmt.Println(err)
	}
	var comment []Comment

	db.First(&comment)
	fmt.Println(comment[0])
}
