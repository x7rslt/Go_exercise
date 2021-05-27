package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"net/http"
	"fmt"
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
	fmt.Println(db)
	if err != nil {
		fmt.Println(err)
	}
	var comment []Comment
	db.Find(&comment)
	fmt.Println(comment)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": comment,
		})
	})
	r.Run()
}

/*
create table comment(
	commentid varchar(64),
	comment text,
	accountname varchar(32),
	accountpic varchar(128),
	star int,
	averageperson int,
	isgood int);

insert into comment(
	commentid,comment,accountname,accountpic,star,averageperson,isgood)
	values ("1","这家饭店还不错，值得尝试，推荐菜养生黑豆腐","乐享生活","food.pic",1299,68,1);
*/
