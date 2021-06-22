package dev_test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"testing"
)

func GetDummyEndpoint(c *gin.Context) {
	resp := map[string]string{"hello": "world"}
	c.JSON(200, resp)
}

func TestGin(t *testing.T) {
	api := gin.Default()
	api.GET("/dummy", GetDummyEndpoint)
	api.Run(":5000")
}

func DummyMiddleware(c *gin.Context) {
	fmt.Println("Im a dummy!")
	c.Next() //调用了c.Next()，这意味着在我们的中间件完成执行后，我们可以将请求处理程序传递给链中的下一个func
}

func TestMiddleWare(t *testing.T) {
	api := gin.Default()
	api.Use(DummyMiddleware)
	api.GET("/middledummy", GetDummyEndpoint)
	api.Run(":80")
}

//API认证中间件

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}
	c.JSON(code, resp)
	c.Abort()
}

func TokenAuthWithMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.FormValue("api_token")
		if token == "" {
			respondWithError(401, "API token required", c)
			return
		}
		if token != os.Getenv("API_TOKEN") {
			respondWithError(401, "Invalid API token", c)
			return
		}
		c.Next()
	}
}

func TestApiAuth(t *testing.T) {
	api := gin.Default()
	api.Use(TokenAuthWithMiddleware())
	api.GET("/apiauth", GetDummyEndpoint)
	api.Run(":80")
}
