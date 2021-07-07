package repository

import (
	"food/model"
	"log"
)

type TeamRepo struct {
	DB model.DataBase
}

func (t *TeamRepo) GetTeamListByHotelId(hotelId string) []model.Team {
	var teamList []model.Team
	t.DB.MyDB.Where("hotel_id=?", hotelId).Find(&teamList)
	log.Println("Repository teamlist hotelid:",hotelId)
	log.Println(teamList)
	return teamList
}

func (t *TeamRepo) GetTeamDetail(teamDetailId string) []model.TeamAggregation {
	var teamAggregations []model.TeamAggregation
	t.DB.MyDB.Table("teamdetail").Where("teamdetail.teamdetailid=?", teamDetailId).Joins("JOIN teamchoosefood on teamdetail.teamdetailid=teamchoosefood.teamdetailid").Joins("JOIN teamchooseitem on teamchoosefood.teamchoosefoodid=teamchooseitem.teamchoosefoodid").Select("teamdetail.*,teamchoosefood.*,teamchooseitem.*").Scan(&teamAggregations)
	return teamAggregations
}
