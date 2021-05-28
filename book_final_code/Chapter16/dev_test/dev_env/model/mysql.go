package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	//"net/http"
	"gorm.io/gorm/schema"
)

type DataBase struct {
	MyDB *gorm.DB
}

var DB *DataBase

func (db *DataBase) Init() {
	DB = &DataBase{MyDB: GetMySqlDB()}
}
func (db *DataBase) Close() {
	DB.MyDB.Close()
}

func GetMySqlDB() *gorm.DB {
	dsn := "root:Xss8271329.X@tcp(167.99.155.35:3306)/food_app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
	}})
	if err != nil {
		log.Println("DataBase connect error:", err)
	}
	return db
}
