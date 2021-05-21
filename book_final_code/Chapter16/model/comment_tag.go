package model

type CommentTag struct {
	TagId   string `json:"tagId"`
	TagName string `json:"tagName"`
	TagNum  int    `json:"tagNum"`
	HotelId string `json:"hotelId"`
}
