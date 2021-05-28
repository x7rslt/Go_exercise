package service

import (
	"dev_env/model"
	"dev_env/repository"
)

type CommentService struct {
	Repo respository.CommentRepo
}

func (c *CommentService) GetCommentList() []model.Comment {
	return Repo.GetCommentList()
}
