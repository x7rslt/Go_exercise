package model

type TeamChooseItem struct {
	TeamChoseItemId   string `gorm:"column:team_chose_item_id"json:"choseItemId"`
	TeamChoseItemName string `gorm:"column:team_chose_item_name"json:"choseItemName"`
	TeamChoseItemTip  string `gorm :"coulumn:team_chose_item_tip"json:"choseItemTip"`
	TeamPrice         int    `gorm:"column:team_price"json:"teamPrice"`
}
