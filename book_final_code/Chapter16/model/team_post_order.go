package model

type TeamPostOrder struct {
	TeamPostOrderId string `json:"teamPostOrderId"`
	TeamDetailId  string `json:"teamDetailId"`
	RealPric      string `json:"realPrice"`
	Quatity       int    `json:"quantity"`
	Mobile        string `json:"mobile"`
	Name          string `json:"name"`
	Sex           int    `json:"sex"`
	Message       string `json:"message"`
}
