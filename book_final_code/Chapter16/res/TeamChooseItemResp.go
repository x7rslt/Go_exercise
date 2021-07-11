package res
type TeamChooseItemResp struct {
	TeamChooseItemId string `gorm:"column:team_choose_item_id" json:"teamChooseItemId"`
	TeamChooseItemName string `gorm:"column:team_choose_item_name" json:"teamChooseItemName"`
	TeamChooseItemTip string `gorm:"column:team_choose_item_tip" json:"teamChooseItemTip"`
	TeamPrice int `gorm:"column:team_price" json:"teamPrice"`
}
