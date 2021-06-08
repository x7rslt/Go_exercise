package repository

import "food/model"

type HotelRepo struct {
	DB model.DataBase
}

func (h *HotelRepo) GetHotelList(hotelid string) model.Hotel {
	var hotel model.Hotel
	h.DB.MyDB.Where("hotelid=?", hotelid).First(&hotel)
	return hotel
}
