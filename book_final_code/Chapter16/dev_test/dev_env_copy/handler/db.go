package handler

import (
	"dev_env/service"
	"fmt"
	//"github.com/gin-gonic/gin"
	//"io/ioutil"
	//"log"
	//"net/http"
	//"os"
)

type CommentHandler struct {
	Srv *service.CommentService
}

func (h *CommentHandler) CommentHandler() {
	fmt.Println("Get Comments.")
	comments := h.Srv.GetCommentList() //这里有问题
	//c.IndentedJSON(200, comments)
	fmt.Println(comments)

}
