package myerr

var (
	OK                  = &ErrNum{Code: 0, Message: "OK"}
	InternalServerError = &ErrNum{Code: 3001, Message: "内部错误"}
	ErrBind             = &ErrNum{Code: 3002, Message: "请求信息无法转换成结构体"}
)
