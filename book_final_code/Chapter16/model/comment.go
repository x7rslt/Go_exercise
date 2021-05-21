package model

type Comment struct {
	CommentId string `json:"commentId"`
	Comment string `json:"comment"`
	AccountName string `json:"accountName"`
	AccountPic string `json:"accountPic"`
	Star int `json:"star"`
	AveragePerson int `json:"averagePerson"`
	IsGood int `json:"isGood"`
}