package model

type Discount struct {
	Id                    string `json:"id"`
	DiscountId            string `json:"discountId"`
	DiscountItemIcon      string `json:"discountItemIcon"`
	DiscountItemSrc       string `json:"discountItemSrc"`
	DiscountItemTitle     string `json:"discountItemTitle"`
	DiscountItemHotel     string `json:"discountItemHotel"`
	DiscountItemGoodPrice int32  `json:"discountItemGoodPrice"`
	DiscountItemPrice     int    `json :dicountItemPrice`
	Discount              string `json:"discount"`
	Pos                   int    `json:"pos"`
}
