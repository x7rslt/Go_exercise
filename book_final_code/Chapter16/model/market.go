package model

type Market struct {
	MarketId   string `json:"marketId"`
	MarketName string `json:"markteName"`
	Addr       string `json:"addr"`
	ShortType  string `json:"shortType"`
	Star       int    `json:"star"`
	Pic        string `json:"pic"`
}
