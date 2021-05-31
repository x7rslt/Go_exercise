package service

import (
	"dev_env/model"
	"dev_env/repository"
)

type CommentService struct {
	Repo repository.CommentRepo
}

func (c *CommentService) GetCommentList() []model.Comment {
	return c.Repo.GetCommentList()
}
