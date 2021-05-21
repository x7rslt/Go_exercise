package model

type Team struct {
	TeamId    string `json:"teamId"`
	HotelId   string `json:"hotelId"`
	TeamName  string `json:"teamName"`
	Price     int    `json:"price"`
	TeamPrice int    `json:"teamPrice"`
	SoldNum   string `json:"soldNum"`
}
