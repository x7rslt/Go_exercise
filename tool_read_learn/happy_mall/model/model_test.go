package model_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Banner struct {
	BannerID    string `json:"bannerID" gorm:"column:banner_id"`
	Url         string `json:"url" gorm:"column:url"`
	RedirectUrl string `json:"redirectUrl" gorm:"column:redirect_url"`
	Order       int    `json:"order" gorm:"column:order_by"`
	CreateUser  string `json:"createUser" gorm:"column:create_user"`
	UpdateUser  string `json:"updateUser" gorm:"column:update_user"`
}

//Json test
type Bannerjson struct {
	BannerID string `json:"bannerID"`
	Url      string `json:"url"`
}

func TestJson(t *testing.T) {
	//Json 主要解析返回数据留作前端展示
	resd := Bannerjson{"1", "http://www.test.com"}
	resJson, _ := json.Marshal(resd)
	fmt.Println(string(resJson))
}

//数据表自动迁移初始化
func TestGormMigrate(t *testing.T) {
	db, _ := gorm.Open("mysql", "root:Xss8271329@tcp(167.99.155.35:3306)/gorm?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	db.SingularTable(true)    //单数格式，例：创建user表，而不是users；
	db.AutoMigrate(&Banner{}) //Gorm mysql 自动迁移初始化，创建Banner表,表字段由gorm:"column:banner_id"设置，column为gorm tag参数；
}

//插入数据
func TestInsert(t *testing.T) {
	db, _ := gorm.Open("mysql", "root:Xss8271329@tcp(167.99.155.35:3306)/gorm?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	db.SingularTable(true)
	ba := Banner{"2", "http://www.test.com", "http://www.test.com", 666, "test", "test"} //Banner赋值
	db.Create(&ba)
	fmt.Println("Insert Done") //Banner表插入数据
}

//查询数据
func TestQuery(t *testing.T) {
	var ba Banner
	var bas []Banner
	db, _ := gorm.Open("mysql", "root:Xss8271329@tcp(167.99.155.35:3306)/gorm?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	db.SingularTable(true)
	db.First(&ba) // Get the first record ordered by primary key = SELECT * FROM users ORDER BY id LIMIT 1;
	fmt.Println(&ba)
	db.Find(&bas) //查询Banner列表 =select * from banner;
	fmt.Println(&bas)
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

func TestGetUserList(t *testing.T) {
	var users []*User
	db, _ := gorm.Open("mysql", "root:Xss8271329@tcp(167.99.155.35:3306)/happy_mall?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	db.SingularTable(true)
	if err := db.Find(&users).Error; err != nil {
		fmt.Println(err)
	}
	fmt.Println(users)
}

//数据库更新
func TestUpdate(t *testing.T) {
	var ba Banner
	var bas []Banner
	db, _ := gorm.Open("mysql", "root:Xss8271329@tcp(167.99.155.35:3306)/gorm?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	db.SingularTable(true)
	//注意是字段列全更新
	//单个字段更新
	db.Model(&ba).Update("url", "http://www.update.com")
	// UPDATE ba SET url='http://www.update.com', updated_at='2013-11-17 21:34:10' WHERE id=111;

	//多字段更新
	// Update attributes with `struct`, will only update non-zero fields
	db.Model(&ba).Updates(Banner{Url: "http://updates.com", BannerID: "111"})
	// UPDATE ba SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;

	db.Find(&bas) //查询Banner列表 =select * from banner;
	fmt.Println(&bas)

}
