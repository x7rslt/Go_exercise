package model

type TeamDetail struct {
	TeamDetailId   string `gorm :"column:team_detail_id"json:"teamDetailId"`
	TeamDetailName string `gorm :"column:team_detail_name"json:"teamDetailName"`
	Policy         string `gorm:"column:policy"json:"policy"`
	Tips           string `gorm:"column:tips"json:"tips"`
}
type TeamAggregation struct {
	TeamDetail
	TeamChooseFood
	TeamChooseItem
}
