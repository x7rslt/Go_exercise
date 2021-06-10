package repository

import (
	"food/model"
)

type SuggestFoodRepo struct {
	DB model.DataBase
}

func (s *SuggestFoodRepo) GetFoodByHotelId(hotelid string) []model.SuggestFood {
	var foods []model.SuggestFood
	s.DB.MyDB.Where("hotelid=?", hotelid).Find(&foods)
	return foods
}

func (s *SuggestFoodRepo) ListSugguest(level int) []model.Suggest {
	var items []model.Suggest
	s.DB.MyDB.Where("level=?", level).Find(&items)
	return items
}
