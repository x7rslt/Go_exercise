package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
)

const(Limit = 20)
func GetRequestID(c *gin.Context)string{
	v,ok:= c.Get("X-Request-Id")
	if !ok{
		return ""
	}
	if requestId,ok:= v.(string);ok{
		return requestId
	}
	return ""

}

func GetMD5Encode(data string)string{
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
