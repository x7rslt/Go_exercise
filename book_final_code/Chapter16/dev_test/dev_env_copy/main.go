package main

import (
	"dev_env/handler"
	"dev_env/model"
	//"github.com/gin-gonic/gin"
	//"log"
	//"net/http"
	"fmt"
)

var (
	CommentHandler handler.CommentHandler
)

func Load() {
	CommentHandler.CommentHandler()
}
func init() {
	initDB()
	initHandler()
	fmt.Println("main init CommentHandler:", CommentHandler)

}

func main() {
	model.DB.Init()
	defer model.DB.Close()
	fmt.Println("CommentHandler:", &CommentHandler)
	Load()
	//port := ":8080"
	//log.Println(http.ListenAndServe(port, r).Error())
}
