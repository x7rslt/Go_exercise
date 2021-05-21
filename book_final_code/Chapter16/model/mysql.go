package model

import (
	"github.com/jinzhu/gorm"
)

type DataBase struct {
	MyDB *gorm.DB
}
