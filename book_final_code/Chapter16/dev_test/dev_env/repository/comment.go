package repository

import (
	"dev_env/model"
	"log"
)

type CommentRepo struct {
	DB model.DataBase
}

func (c *CommentRepo) GetCommentList() model.Comment {
	fmt.Println("Repository no probelom")
	var commentList []model.Comment
	c.DB.MyDB.Find(&commentList)
	log.Println(commentList)
	return commentList[0]
}
