package repository

import (
	"food/model"
)

type RestaurantNavRepo struct {
	DB model.DataBase
}

func (r *RestaurantNavRepo) ListRestaurantNav(level int) []model.RestaurantNav {
	var items []model.RestaurantNav
	r.DB.MyDB.Where("level=?", level).Find(&items)
	return items
}

type RestaurantTabItemRepo struct {
	DB model.DataBase
}

func (r *RestaurantTabItemRepo) ListGoodRestaurantTabItem() []model.RestaurantTableItem {
	var items []model.RestaurantTableItem
	r.DB.MyDB.Find(&items)
	return items
}

func (r *RestaurantTabItemRepo) ListRestaurantNav(level int) []model.RestaurantNav {
	var items []model.RestaurantNav
	r.DB.MyDB.Where("level=?", level).Find(&items)
	return items
}
