package res

import "food/model"

type TeamResp struct{
	model.TeamDetail
	FoodList []TeamChooseFoodResp
	ChooseList []TeamChooseItemResp
}
