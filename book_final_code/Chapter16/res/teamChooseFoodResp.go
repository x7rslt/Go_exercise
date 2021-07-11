package res

type TeamChooseFoodResp struct {
	TeamChooseFoodId string `gorm:"column:team_choose_food_id" json:"teamChooseFoodId"`
	TeamChooseFoodName string `gorm:"column:team_choose_food_name" json:"teamChooseFoodName"`
}