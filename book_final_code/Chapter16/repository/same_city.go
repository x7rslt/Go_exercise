package repository

import (
	"food/model"
)

type SameCityRepo struct {
	DB model.DataBase
}

func (s *SameCityRepo) ListSameCity() []model.SameCity {
	var samecity []model.SameCity
	s.DB.MyDB.Find(&samecity)
	return samecity
}
