package repository

import (
	"food/model"
)

type ListNavItemRepo struct {
	DB model.DataBase
}

func (n *ListNavItemRepo) ListNavItems(level int) []model.Nav {
	var items []model.Nav
	n.DB.MyDB.Where("level=?", level).Find(&items)
	return items
}
