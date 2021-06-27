package handler

import (
	"food/service/dp_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

const domain = "http://192.168.0.104:8080"

type HotelDetailHandler struct {
	Srv *dp_service.HotelService
}

func (h *HotelDetailHandler) HotelDetailHandler(c *gin.Context) {
	hotelId := c.Param("id")

	hotel, err := h.Srv.GetHotelDetailByID(hotelId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"item": nil,
			"msg":  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"item": hotel,
			"msg":  "",
		})
	}
}
