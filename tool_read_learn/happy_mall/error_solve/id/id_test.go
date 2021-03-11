package id_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ListQuery struct {
	PageSize int `json:"pageSize"`
	Page     int `json:"page"`
}

type UserRepository struct {
	DB *gorm.DB
}

// Page 分页
func Page(Limit, Page int) (limit, offset int) {
	if Limit > 0 {
		limit = Limit
	} else {
		limit = 10
	}
	if Page > 0 {
		offset = (Page - 1) * limit
	} else {
		offset = -1
	}
	return limit, offset
}

func (repo *UserRepository) List(req *ListQuery) (users []*User, err error) {
	fmt.Println(req)
	db := repo.DB
	limit, offset := Page(req.PageSize, req.Page) // 分页

	if err := db.Order("id desc").Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

type User struct {
	UserId    string    `json:"userId" gorm:"column:user_id"`
	NickName  string    `json:"nickName" gorm:"column:nick_name"`
	Mobile    string    `json:"mobile" gorm:"column:mobile" binding:"required"`
	Password  string    `json:"password" gorm:"column:password"`
	Address   string    `json:"address" gorm:"column:address"`
	IsDeleted bool      `json:"isDeleted" gorm:"column:is_deleted"`
	IsLocked  bool      `json:"isLocked" gorm:"column:is_locked"`
	CreateAt  time.Time `json:"createAt" gorm:"column:create_at;default:null"`
	UpdateAt  time.Time `json:"updateAt" gorm:"column:update_at;default:null"`
}

func Test(t *testing.T) {
	var users []*User
	db, _ := gorm.Open("mysql", "root:Xss8271329.X@tcp(167.99.155.35:3306)/happy_mall?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	req := ListQuery{1, 2}
	limit, offset := Page(req.PageSize, req.Page) // 分页
	fmt.Println(limit, offset)
	db.SingularTable(true)
	if err := db.Order("user_id desc").Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		fmt.Println(nil, err)
	}
	fmt.Println(users, nil)

}
