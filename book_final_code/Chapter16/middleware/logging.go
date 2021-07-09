package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"food/MyLog"
	"food/model/log"
	"food/myerr"
	"food/res"
	"github.com/willf/pad"
	"io/ioutil"
	"strings"
	"time"
)

func Logging() gin.HandlerFunc{
	return func(context *gin.Context){
		start := time.Now().UTC()
		path := context.Request.URL.Path

		requrl := context.Request.URL.Path
		if strings.Contains(requrl,"image"){
			return
		}
		var bodies []byte
		if context.Request.Body != nil{
			bodies,_=ioutil.ReadAll(context.Request.Body)


		}
		context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodies))
		method := context.Request.Method
		ip := context.ClientIP()

		blw := &log.BodyLoggerWriter{
			ResponseWriter:context.Writer,
			Body:bytes.NewBufferString(""),

		}
		context.Writer= blw
		end := time.Now().UTC()
		sub := end.Sub(start)
		code := -100
		message := ""
		var response res.Response
		if err := json.Unmarshal(blw.Body.Bytes(),&response);err!=nil{
			code = myerr.InternalServerError.Code
			message = err.Error()
			MyLog.LogError(err)

		}else{
			code = response.Code
			message = response.Message

		}
		MyLog.Log.Infof("%-13s |%-12s|%s %s|{code:%d,message:%s}",sub,ip,pad.Right(method,5,""),path,code,message)

	}
}