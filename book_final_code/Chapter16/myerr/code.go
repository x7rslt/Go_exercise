package myerr

var (
	//Common errors
	OK = &ErrNum{Code:0,Message: "OK"}
	InternalServerError= &ErrNum{Code: 30001,Message: "内部错误"}
	ErrBind=&ErrNum{Code: 30002,Message: "请求信息无法转换成结构体"}
	ErrDatabase = &ErrNum{30002,"Database error"}
	ErrValidation = &ErrNum{30001,"Validation failed"}
	ErrEncrypt = &ErrNum{30101,"Error occurred while encryption the user password"}
	ErrAccountNotFound = &ErrNum{50102,"用户不存在"}
	ErrPassword = &ErrNum{50103,"密码错误"}
	ErrAccountEmpty = &ErrNum{50104,"用户名不能为空"}
	ErrPasswordEmpty = &ErrNum{50103,"密码不能为空"}
	ErrMissingHeader = &ErrNum{50104,"Header不存在"}
	ErrToken = &ErrNum{50105,"生成Token错误"}
	PassParamCheck = &ErrNum{60000,"参数校验通过"}
)