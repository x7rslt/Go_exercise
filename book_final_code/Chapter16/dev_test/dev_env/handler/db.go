package handler

import (
	"dev_env/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	//"net/http"
	"os"
)

type CommentHandler struct {
	Srv *service.CommentService
}

func (h *CommentHandler) CommentHandler(c *gin.Context) {
	fmt.Println("Get Comments.")
	comments := h.Srv.GetCommentList() //这里有问题
	c.IndentedJSON(200, comments)

}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func JsonHandler(c *gin.Context) {
	//JSON 美化输出
	xiao := User{0, "xss", 28}
	wang := User{1, "wang", 30}
	zhang := User{2, "zhang", 28}
	c.IndentedJSON(200, []User{xiao, wang, zhang})
}

func ImageHandler(c *gin.Context) {
	imageName := c.Query("imageName")
	fmt.Println(imageName)
	dir, _ := os.Getwd()
	file, err := ioutil.ReadFile(dir + "/static/images/" + imageName + ".png")
	if err != nil {
		log.Println(err)
	}
	c.Writer.WriteString(string(file))
}
