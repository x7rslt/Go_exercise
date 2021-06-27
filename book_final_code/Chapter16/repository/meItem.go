package repository

import (
	"food/model"
)

type MeItemRepo struct {
	DB model.DataBase
}

func (item *MeItemRepo) ListMe() []model.MeItem {
	var items []model.MeItem
	item.DB.MyDB.Find(&items)
	return items
}
