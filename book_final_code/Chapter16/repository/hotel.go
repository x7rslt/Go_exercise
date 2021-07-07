package repository

import "food/model"

type HotelRepo struct {
	DB model.DataBase
}

func (h *HotelRepo) GetHotelById(hotelid string) model.Hotel {
	var hotel model.Hotel
	h.DB.MyDB.Where("hotel_id=?", hotelid).First(&hotel)
	return hotel
}
