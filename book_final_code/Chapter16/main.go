package main

import (
	"fmt"
	"food/model"
	"github.com/spf13/viper"
)

func main() {
	test := new(model.Hotel)
	fmt.Println(test)
	parameter := viper.GetString("database.username")
	fmt.Println(parameter)
}
