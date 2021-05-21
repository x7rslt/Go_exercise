package model

type TakeOut struct {
	TakeOutId           string `json:"takeOutId"`
	HotelFoodCategoryId string `json:"hotelFoodCategoryId"`
	FoodName            string `json:"foodName"`
	Pic                 string `json:"pic"`
	MonthSoldNum        string `json:"monthSoldNum"`
	Zan                 int    `json :"zan"`
	Pric                int    `json:"price"`
	IsSuggest           int    `json:"isSuggest"`
	DiscountPrice       int    `json:"discountPrice"`
	Discount            int    `json:"discount"`
}

type TakeCount struct {
	Hotel
	HotelFoodCategory
	TakeOut
}
