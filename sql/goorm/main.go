package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	//dsn := "root:Xss8271329@tcp(204.44.98.24:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//Sqlite
	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db)
}
