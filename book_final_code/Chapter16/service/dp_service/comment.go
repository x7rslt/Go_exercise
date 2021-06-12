package dp_service

import (
	"food/model"
	"food/repository"
)

type CommentService struct {
	Repo repository.CommentRepo
}

func (c CommentService) GetCommentList(hotelId string) []model.Comment {
	return c.Repo.GetCommentList(hotelId)
}
