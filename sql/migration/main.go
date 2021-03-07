package main

import (
	"fmt"
	"time"

	"Go_exercise/sql/migration/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, _ := gorm.Open("mysql", "root:Xss8271329.X@tcp(167.99.155.35:3306)/happy_mall?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	db.SingularTable(true)
	//db.AutoMigrate(&model.Banner{}, &model.Category{}, &model.OrderItem{}, &model.Order{}, &model.Product{}, &model.ShoppingCartItem{}, &model.Stock{}, &model.User{})*/
	user1 := model.User{"2", "test2", "15555", "123456", "hangzhou", false, false, time.Now(), time.Now()}
	db.Create(&user1)
	fmt.Println(db)
}
