package main

import (
	"Project/book_final_code/Chapter17/router"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	G := gin.New()
	middlewares := []gin.HandlerFunc{}
	router.Load(
		g,
		middlewares...,
	)
	go func() {
		if err := CheckServer(); err != nil {
			log.Fatal("自检程序发生错误...", err)
		}
		log.Println(路由部署成功)
	}()
	log.Printf("开始监听http地址：%s", "9090")
	log.Printf(http.ListenAndServe(":9090", g).Error())
}

func CheckServer() error {
	for i := 0; i < 10; i++ {
		resp, err := http.Get("http://127.0.0.1:9090/check/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		log.Print("等待路由，1秒后重试")
		time.Sleep(1 * time.Second)
	}
	return errors.New("无法连接路由")
}
