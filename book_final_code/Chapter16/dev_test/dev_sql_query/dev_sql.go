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
	CommentId string `json:"commentId"`
	//评论
	Comment string `json:"comment"`
	//评论人
	AccountName string `json:"accountName"`
	//评论人头像
	AccountPic string `json:"accountPic"`
	//评分
	Star int `json:"star"`
	//人均消费
	AveragePerson int `json:"averagePerson"`
	//是否是优质点评
	IsGood int `json:"isGood"`
}

func (Comment) TablesName() string {
	return "comment"
}
func main() {

	dsn := "root:Xss8271329.X@tcp(167.99.155.35:3306)/food_app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Printf("%v", &schema.Namer{})
	fmt.Println(db)
	if err != nil {
		fmt.Println(err)
	}
	var comment []Comment

	db.Find(&comment)
	fmt.Println(comment)
}
