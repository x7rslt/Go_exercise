package dev_test

import (
	"fmt"
	"food/model"
	"testing"
)

//Comment Repository
type CommentRepo struct {
	DB model.DataBase
}

func (c *CommentRepo) GetCommentList(hotelId string) []model.Comment {
	var commentList []model.Comment
	c.DB.MyDB.Where("hotel_id=?", hotelId).Find(&commentList)
	return commentList
}

//Comment Service
type CommentService struct {
	Repo CommentRepo
}

func (c *CommentService) GetCommentList(hotelId string) []model.Comment {
	return c.Repo.GetCommentList(hotelId)
}

//Comment Handler
const domain = "http://192.168.0.104:8080"

type CommentHandler struct {
	Srv *dp_service.CommentService
}

func (h *CommentHandler) CommentList(c **gin.Context) {
	hotelId := c.Param("id")
	comments := h.Src.GetCommentList(hotelId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"itme": nil,
			"msg":  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"item": comments,
			"msg":  "",
		})
	}

}

func TestDev(t *testing.T) {
	test := new(model.Hotel)
	fmt.Println(test)
	db := model.DataBase{}
	fmt.Println(db)
}
