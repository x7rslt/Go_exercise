package main

import (
	"dev_env/handler"
	"dev_env/health"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(engine *gin.Engine) *gin.Engine {
	engine.Use(gin.Recovery())
	engine.NoRoute(func(context *gin.Context) {
		context.String(http.StatusNotFound, "API路由不正确")
	})
	check := engine.Group("/check")
	{
		check.GET("/health", health.Health)
	}
	dp := engine.Group("/v1/dp")
	{
		dp.GET("/comments", CommentHandler.CommentHandler)
		//获取图片
		dp.GET("/image", handler.ImageHandler)
		dp.GET("/json", handler.JsonHandler)

	}
	return engine
}
