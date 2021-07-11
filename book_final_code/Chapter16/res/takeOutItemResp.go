package res

type TakeOutItemResp struct{
	TakeOutId string `json:"takeOutId"`
	HotelFoodCategoryId string `json:"hotelFoodCategoryId"`
	FoodName  string `json:"foodName"`
	Pic string `json:"pic"`
	MonthSoldNum string `json:"monthSoldNum"`
	Zan int `json:"zan"`
	Price int `json:"price"`
	IsSuggest int `json:"isSuggest"`
	DiscountPrice int `json:"discountPrice"`
	Discount float64 `json:"discount"`

}