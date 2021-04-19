package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	s := gin.Default()
	//Query
	s.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, c.Query("order"))
	})
	//DefaultQuery
	s.GET("/2", func(c *gin.Context) {
		c.String(http.StatusOK, c.DefaultQuery("id", "0"))
	})
	//QueryArray
	s.GET("/3", func(c *gin.Context) {
		c.JSON(200, c.QueryArray("media"))
	})
	//QueryMap
	s.GET("/4", func(c *gin.Context) {
		c.JSON(200, c.QueryMap("ids"))
	})
	//Form
	s.POST("/5", func(c *gin.Context) {
		name := c.PostForm("name")
		c.String(200, "hello "+name)
	})
	//JSON
	xiao := User{0, "xss", 28}
	wang := User{1, "wang", 30}
	zhang := User{2, "zhang", 28}
	s.GET("/users", func(c *gin.Context) {
		c.IndentedJSON(200, []User{xiao, wang, zhang})
	})
	s.Run(":80")
}
