package repository_test

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

func (repo *UserRepository) Exist(user User) int {
	var u User
	repo.DB.Where("nick_name = ?", user.NickName).Find(&u)
	if u.UserId != "" {
		fmt.Println("User exist!")
		return 1
	}
	return 0
}

func (repo *UserRepository) Add(user User) error {
	if repo.Exist(user) == 1 {
		return fmt.Errorf("用户已存在")
	}
	error := repo.DB.Create(&user).Error
	if error != nil {
		return fmt.Errorf("用户注册失败")
	}
	fmt.Println("用户注册成功！")
	return nil
}

func TestGetUserList(t *testing.T) {
	var users []User
	db, _ := gorm.Open("mysql", "root:Xss8271329@tcp(167.99.155.35:3306)/happy_mall?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	db.SingularTable(true)
	if err := db.Find(&users).Error; err != nil {
		fmt.Println(err)
	}
	fmt.Println(users)
}

func TestUserContrl(t *testing.T) {
	db, _ := gorm.Open("mysql", "root:Xss8271329@tcp(167.99.155.35:3306)/happy_mall?charset=utf8&parseTime=True&loc=Local")
	d := UserRepository{db}
	newuser := User{"7", "x7rslt2", "10000000000", "12345678", "hangzho", false, false, time.Now(), time.Now()}
	d.DB.SingularTable(true)
	result := d.Exist(newuser)
	fmt.Println(result)

	a := d.Add(newuser)
	fmt.Println(a)

}

func TestWhere(t *testing.T) {
	db, _ := gorm.Open("mysql", "root:Xss8271329@tcp(167.99.155.35:3306)/happy_mall?charset=utf8&parseTime=True&loc=Local")
	db.SingularTable(true)

	var user1 User
	db.Find(&user1)
	fmt.Println("user1", user1) //注意：查询结果赋值到user1，且nick_name不存在状态下也会返回默认值

	var user2 User
	db.Where("nick_name = ?", "x7rslt").Find(&user2) //注意Where和find的顺序，Where在前查询nick——name，find在前之间赋值第一条数据到user2，where条件不生效。
	// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;
	//fmt.Println(result)
	fmt.Println("user2", user2) //注意：查询结果赋值到user1，且nick_name不存在状态下也会返回默认值
	if user2.UserId == "" {
		fmt.Println("User not exist!")
	}

}
