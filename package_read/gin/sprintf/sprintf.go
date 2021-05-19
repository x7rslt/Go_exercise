package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//query + post form
	router.POST("/post", func(c *gin.Context) {

		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		msg := fmt.Sprintf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
		c.String(200, msg)
	})
	router.Run(":8080")
}
