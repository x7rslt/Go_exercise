package handler

import (
	"dev_env/repository"
	"dev_env/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"io/ioutil"
)

const domain = "http://192.168.0.104:8080"

type CommentHandler struc {
	Srv *service.CommentService
}

func (h *CommentHandler) CommentList(c *gin.Context){
	comments :=h.Srv.GetCommentList()
	c.JSON(http.StatusOK,gin.H{
		"items":comments,
	}) 
}

func ImageHandler(c *gin.Context){
	imageName := c.Query("imageName")
	fmt.Println(imageName)
	dir,_:= os.Getwd()
	file,err := ioutil.ReadFile(dir +"/static/images/"+imageName+".png")
	if err!=nil{
		log.Println(err)
	}
	c.Write.WriteString(string(file))
}