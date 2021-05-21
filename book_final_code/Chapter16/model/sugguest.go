package model

type Suggest struct {
	SuggestId string `json:"suggestId"`
	KeyWord1  string `	json:"keyWord1"`
	KeyWord2  string `json:"keyWord2"`
	KeyWord3  string `json:"keyWord3"`
	KeyWord4  string `json:"keyWord4"`

	Src         string `json:"src"`
	FoodName    string `json:"foodName"`
	Icon        string `json:"icon"`
	HotelName   string `json:"hotelName"`
	JoinPersons string `json:"joinPersons"`
	Price       string `json:"price"`
	GoodPrice   string `json:"goodPrice"`
	Distance    string `json:"distance"`
	Level       int    `json:"level"`
}
