package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, _ := gorm.Open("mysql", "root:Xss8271329.X@tcp(167.99.155.35:3306)/happy_mall?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	db.SingularTable(true)
	//db.AutoMigrate(&model.Banner{}, &model.Category{}, &model.OrderItem{}, &model.Order{}, &model.Product{}, &model.ShoppingCartItem{}, &model.Stock{}, &model.User{})*/
	/*ba := model.Banner{"1", "http://www.baidu.com", "http://www.x7rslt.com", 666, "test", "test"}
	ca := model.Category{"1", "类目", "desc", 10010, "1", false}
	oi := model.OrderItem{"Apple", "1", "1", "hongfushi", "http://product.com", 10, 100, "test", "test"}
	or := model.Order{"1", "1", "custom", "15555", 1000, 1, 0, "20210307", 1, "no pay", "hangzhou", false, "test", "test"}
	pr := model.Product{"1", "香蕉", "haochi", "1", "http://www.baidu.com", "1", 10, 11, 100, "yonghui", 1, "test", "test", "aimaibumai", false}
	sc := model.ShoppingCartItem{"1", "1", "1", 1, false, "tset", "test"}
	st := model.Stock{"1", 100}
	db.Create(&ba)
	db.Create(&ca)
	db.Create(&oi)
	db.Create(&or)
	db.Create(&pr)
	db.Create(&sc)
	db.Create(&st)
	fmt.Println(db)*/
}
