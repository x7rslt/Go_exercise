package respository

import "dev_env/model"

type CommentRepo struct {
	DB model.DataBase
}

func (c *CommentRepo) GetCommentList() {
	var commentList []mode.Comment
	c.DB.MyDB.Find(&commentList)
	return commentList
}
