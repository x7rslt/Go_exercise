package service_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

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
type UserRepository struct {
	DB *gorm.DB
}

type UserRepoInterface interface {
	List(req *ListQuery) (users []*User, err error)
}

func (repo *UserRepository) List(req *ListQuery) (users []*User, err error) {
	fmt.Println(req)
	db := repo.DB
	if err := db.Order("user_id desc").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

type UserSrv interface {
	List(req *ListQuery) (users []*User, err error)
}

type UserService struct {
	Repo UserRepoInterface
}

func (srv *UserService) List(req *ListQuery) (users []*User, err error) {
	if req.PageSize < 1 {
		req.PageSize = 0
	}
	return srv.Repo.List(req)
}

type ListQuery struct {
	PageSize int
	Page     int
}

func TestService(t *testing.T) {
	db, _ := gorm.Open("mysql", "root:Xss8271329@tcp(167.99.155.35:3306)/happy_mall?charset=utf8&parseTime=True&loc=Local")
	d := &UserRepository{db}
	d.DB.SingularTable(true)
	repo := UserService{d}
	query := ListQuery{1, 1}
	result, err := repo.List(&query)
	fmt.Println(*result[1], err)
	for _, i := range result {
		fmt.Println(*i)
	}
}
