package main

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	Driver = "mysql"
	DsName = "root:Xss8271329@tcp(204.44.98.24:3306)/gotest?charset=utf8" //数据库需提前创建
	DB     *xorm.Engine
	DBErr  error
)

type User struct {
	Id       int64     `xorm:"pk autoincr bigint(64)" form:"id" json:"id"`
	Mobile   string    `xorm:"varchar(20)" form:"mobile" json:"mobile"`
	Password string    `xorm:"varchar(40)" form:"password" json:"-"`
	Avatar   string    `xorm:"varchar(150)" form:"avatar" json:"avatar"`
	Sex      string    `xorm:"varchar(2)" form:"sex" json:"sex"`
	Nickname string    `xorm:"varchar(20)" form:"nickname" json:"nickname"`
	Salt     string    `xorm:"varchar(10)" form:"salt" json:"-"`
	Online   int       `xorm:"int(10)" form:"online" json:"online"` //是否在线
	Token    string    `xorm:"varchar(40)" form:"token" json:"token"`
	Memo     string    `xorm:"varchar(140)" form:"memo" json:"memo"`
	Createat time.Time `xorm:"datetime" form:"createat" json:"createat"`
}

func main() {
	DB, DBErr = xorm.NewEngine(Driver, DsName)
	if DBErr != nil {
		log.Fatal(DBErr)
	}
	//DB.ShowSQL(true)
	DB.SetMaxOpenConns(2)
	_ = DB.Sync2(
		new(User),
	)
	log.Println("init database success")
}
