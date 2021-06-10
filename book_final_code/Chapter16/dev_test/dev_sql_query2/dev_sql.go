package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	//"net/http"

	"gorm.io/gorm/schema"
)

type User struct {
	gorm.Model
	Name      string
	CompanyID int
	Company   Company `gorm:"foreignKey:CompanyID"`
}

type Company struct {
	ID   int
	Name string
}

func main() {

	dsn := "root:Xss8271329.X@tcp(167.99.155.35:3306)/food_app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
	}})
	fmt.Println(db)
	if err != nil {
		fmt.Println(err)
	}
	//companyid := 1
	var UserList []User
	db.Find(&UserList)
	fmt.Println(UserList)
}
