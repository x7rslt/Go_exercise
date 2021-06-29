package main

import (
	"food/health"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Load(engine *gin.Engine, middlewares ...gin.HandlerFunc) *gin.Engine {
	engine.Use(gin.Recovery())
	engine.Use(middlewares...)
	engine.NoRoute(func(context *gin.Context) {
		context.String(http.StatusNotFound, "API路由不正确.")
	})
	check := engine.Group("/check")
	{
		check.GET("health", health.Health)
	}
	dp := engine.Group("/v1/dp")
	{
		//dp.GET("/index",handler.IndexHandler)

		dp.GET("/hotel/detail/:id", HotelDetailHandler.HotelDetailHandler)
	}
	return engine
}
