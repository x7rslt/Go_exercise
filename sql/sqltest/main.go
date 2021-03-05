package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

/*CREATE TABLE IF NOT EXISTS `users`(
`user_id` VARCHAR(40),
`nick_name` VARCHAR(40),
`mobile` VARCHAR(40),
`Password` VARCHAR(40),
`address` VARCHAR(40),
`is_deleted` bool,
`is_locked` bool,
`create_at` DATE,
`update_at` DATE);*/
type UserRepoInterface interface {
	Get(user User) (*User, error)
	Add(user User) (*User, error)
}

type UserRepository struct {
	DB *gorm.DB
}

func Md5(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}

func (repo *UserRepository) Get(user User) (*User, error) {
	if err := repo.DB.Where(&user).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (repo *UserRepository) Add(user User) {
	repo.DB.Create(&user)
	fmt.Println("Done")
}

//初始化配置
func initConfig() {
	viper.AddConfigPath(`C:\Users\Dell\go\src\happy_mall_backend\sqltest\conf`)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml") // 设置配置文件格式为YAML
	viper.ReadInConfig()
}

//初始化数据库
var DB *gorm.DB

type DBConf struct {
	Driver string
	// Host 主机地址
	Host string
	// Port 主机端口
	Port string
	// User 用户名
	User string
	// Password 密码
	Password string
	// DbName 数据库名称
	DbName string
	// Charset 数据库编码
	Charset string
}

func initDB() *gorm.DB {
	fmt.Println("Database init:")
	var err error
	conf := &DBConf{
		Host:     viper.GetString("database.host"),
		User:     viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		DbName:   viper.GetString("database.name"),
	}
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8&parseTime=%t&loc=%s",
		conf.User,
		conf.Password,
		conf.Host,
		conf.DbName,
		true,
		"Local")
	DB, err = gorm.Open(mysql.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatalf("connect error: %v\n", err)
	}
	//DB.SingularTable(true)
	fmt.Println("数据库 init 结束...")
	return DB
}

func main() {
	initConfig()
	fmt.Println(viper.AllSettings())

	user1 := User{"1", "test", "15555", "123456", "hangzhou", false, false, time.Now(), time.Now()}
	data := initDB()
	d := UserRepository{data}
	d.Add(user1)

}
