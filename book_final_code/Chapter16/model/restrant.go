package model

type RestaurantNav struct {
	NavId string `json:"navId"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
	Src   string `json:"src"`
	Level int    `json:"level"`
}

type RestaurantTableItem struct {
	ItemId string  `json:"itemId"`
	Src    string  `json:"src"`
	Title  string  `json:"title"`
	Star   float32 `json:"star"`
	Pric   string  `json:"price"`
	Area   string  `json:"ar"`
	Type   string  `json:"type"`
	Desc   string  `json:"desc"`
	Team   string  `json:"team"`
	Quan   string  `json:"quan"`
}
