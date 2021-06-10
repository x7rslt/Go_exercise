package repository

import (
	"food/model"
)

type SugguestRepo struct {
	DB model.DataBase
}

func (s *SugguestRepo) ListSugguest(level int) model.Suggest {
	var item model.Suggest
	s.DB.MyDB.Where("level=?", level).Find(&item)
	return item
}
