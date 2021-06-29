package main

import (
	"flag"
	"food/MyLog"
	"food/config"
	"food/handler"
	"food/model"
	"github.com/gin-gonic/gin"
	"github.com/i-coder-robot/book_final_code/Chapter16/middleware"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

var (
	cfg                = flag.String("config", "", "")
	HotelDetailHandler handler.HotelDetailHandler
)

func init() {
	initViper()
	initDB()
	initHandler()
}

func main() {
	flag.Parse()
	if err := config.Init(*cfg);err!=nil{
		panic(err)
	}
	model.DB.Init()
	defer model.DB.Close()
	r := gin.New()
	Load(
		r,middleware.ProcessTraceID(),
		middleware.Logging(),
	)
	port := viper.GetString("addr")
	MyLog.Log.Info("开始监听端口：",port)
	log.Printf(http.ListenAndServe(port,r).Error())
}
