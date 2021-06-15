package main

import (
	"fmt"
	"food/handler"
	"food/model"

	"github.com/spf13/viper"
)

var (
	HotelDetailHandler handler.HotelDetailHandler
)

func main() {
	test := new(model.Hotel)
	fmt.Println(test)
	parameter := viper.GetString("database.username")
	fmt.Println(parameter)
}
