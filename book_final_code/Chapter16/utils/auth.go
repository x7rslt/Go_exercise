package utils

import (
	"github.com/go-playground/validator/v10"
	"food/myerr"

	"food/model"
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(source string)(string,error){
	hashedBytes,err := bcrypt.GenerateFromPassword([]byte(source),bcrypt.DefaultCost)
	return string(hashedBytes),err
}
func Compare(hashedPassword,password string)error{
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password))

}
func Validate(m model.Account)error{
	validate := validator.New()
	return validate.Struct(m)

}
func CheckParam(accountName,password string)myerr.Err{
	if accountName == ""{
		return myerr.New(*myerr.ErrValidation,nil).Add("用户名位空")
		if password ==""{
			return myerr.New(*myerr.ErrValidation,nil).Add("密码为空")

		}
		return myerr.Err(ErrNum:*myerr.PassParamCheck,Err:nil)

	}
}