package model

type Guess struct {
	GuessId   string `json:"guessId"`
	Src       string `json:"src"`
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	GoodPrice string `json:"price"`
	Price     string `json:"price"`
	SoldNum   int32  `json:"soldNum"`
}
