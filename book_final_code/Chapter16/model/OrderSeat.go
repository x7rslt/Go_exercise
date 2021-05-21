package model

type OrderSeat struct {
	OrderSeat string `json:"orderSeatId"`
	HotelID   string `json:"hotelId"`
	Persons   int    `json:"persons"`
	DateTime  int64  `json:"datetime"`
	Mobile    int    `json:"mobile"`
	RoomType  int    `json:"room_Type"`
	Name      string `json:"name"`
	Sex       int    `json:"sex"`
	Message   string `json:"message"`
}
