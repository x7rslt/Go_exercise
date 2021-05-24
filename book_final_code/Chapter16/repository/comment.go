package repository

import "food/model"

type CommentRepo struct {
	DB model.DataBase
}

func (c *CommentRepo) GetCommentList(hoteId string) []model.Comment {
	var commentList []model.Comment
	c.DB.MyDB.Where("hotel_id=?", hotelId).Find(&commentList)
	return commentList
}
