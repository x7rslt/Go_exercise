package service

import (
	"dev_env/model"
	"dev_env/repository"
	"fmt"
)

type CommentService struct {
	Repo *repository.CommentRepo
}

func (c *CommentService) GetCommentList() []model.Comment {
	fmt.Println("Service wrong.")
	return c.Repo.GetCommentList()
}
