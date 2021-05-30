package repository

import (
	"dev_env/model"
	"fmt"
	"log"
)

type CommentRepo struct {
	DB model.DataBase
}

func (c *CommentRepo) GetCommentList() []model.Comment {
	fmt.Println("Repository start.")
	var comment1 *model.Comment
	fmt.Println("Find 1 comment:", comment1)
	c.DB.MyDB.First(&comment1)
	fmt.Println("Find 1 comment:", comment1)
	var commentList []model.Comment
	c.DB.MyDB.Find(&commentList)
	log.Println("Repositroy result:", commentList)
	fmt.Println("Repository no probelom")
	return commentList
}
