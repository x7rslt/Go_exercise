package res
type WXLoginResponse struct {
	OpenId string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId string `json:"unionId"`
	ErrCode int `json:"errCode"`
	ErrMsg string `json:"errmsg"`
}
