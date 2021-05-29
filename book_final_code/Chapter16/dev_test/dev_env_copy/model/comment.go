package model

type Comment struct {
	//评论ID
	Commentid string `json:"commentid"`
	//评论
	Comment string `json:"comment"`
	//评论人
	Accountname string `json:"accountname"`
	//评论人头像
	Accountpic string `json:"accountpic"`
	//评分
	Star int `json:"star"`
	//人均消费
	Averageperson int `json:"averageperson"`
	//是否是优质点评
	Isgood int `json:"isgood"`
}
