package model

type Hotel struct {
	HotelId        string        `json:"hotelId"`
	HotelName      string        `json:"hotelName"`
	Pic            string        `json:"pic"`
	Star           string        `json:"star"`
	Num            int           `json:"num"`
	AveragePrice   float64       `json:"averagePrice"`
	Taste          float64       `json:"taste"`
	Env            float64       `json:"env"`
	Service        float64       `json:"service"`
	MapAddr        string        `json:"mapAddr"`
	MapAddr2       string        `json:"mapAddr2"`
	ShortType      string        `json:"shortType"`
	BusinessTime   string        `json:"businessTime"`
	Bang           string        `json:"bang"`
	TeamList       []Team        `json:"teamList"`
	FoodList       []SuggestFood `json:"foodList"`
	CommentTagList []CommentTag  `json:"commentTagList"`
	CommentList    []Comment     `json:"commentList"`
	Market         Market        `json:"market"`
}
