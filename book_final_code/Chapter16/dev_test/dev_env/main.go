package main

import (
	"dev_env/handler"
	"dev_env/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	CommentHandler handler.CommentHandler
)

func init() {
	initHandler()
	initDB()
}

func main() {
	model.DB.Init()
	defer model.DB.Close()
	r := gin.New()
	Load(r)
	port := ":8080"
	log.Println(http.ListenAndServe(port, r).Error())
}
