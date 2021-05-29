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

	var commentList []model.Comment
	c.DB.MyDB.Find(&commentList)
	log.Println("Repositroy result:", commentList)
	fmt.Println("Repository no probelom")
	return commentList
}
