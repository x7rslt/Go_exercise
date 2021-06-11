package dp_service

import (
	"food/model"
	"food/repository"
)

type CommentTagService struct {
	Repo repository.CommentTagRepo
}

func (c *CommentTagService) GetCommentTaGList(hotelId string) []model.CommentTag {
	return c.Repo.GetCommentTagList(hotelId)
}
