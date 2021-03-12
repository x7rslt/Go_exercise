package gin_test

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGin(t *testing.T) {
	r := gin.Default()
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	r.Run()
}
