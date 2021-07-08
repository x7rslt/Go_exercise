package dev_sql_connect_t_test

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"testing"
	// MySQL driver.
	"food/model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	username string = "root"
	addr string ="114.116.230.93:3306"
	name string ="food"
	password string = "mysqltest110"
)

func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"Local",
	)
	fmt.Println("Model Mysql Config:",config)
	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Printf("DataBase connnect failed.DataBase name:%s,Error:%s", name, err.Error())
	}
	setupDB(db)
	db.SingularTable(true)
	return db

}
func setupDB(db *gorm.DB) {
	db.DB().SetMaxIdleConns(5)

}

func TestMysqlConnect(t *testing.T){
	db:= openDB(username,password,addr,name)
	var teamList []model.Team
	db.Where("hotel_id=?", 1).Find(&teamList)
	log.Println("Teamlist:",teamList)

}