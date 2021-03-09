package utils_test

import (
	"crypto/md5"
	"fmt"
	"io"
	"testing"
	time "time"
)

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

func TestPage(t *testing.T) {
	limit, offset := Page(1, 1)
	fmt.Println(limit, offset)
}

// Sort 排序
// 默认 created_at desc
func Sort(Sort string) (sort string) {
	if Sort != "" {
		sort = Sort
	} else {
		sort = "create_at desc"
	}
	return sort
}
func TestSort(t *testing.T) {
	sort := Sort("test")
	fmt.Println(sort)
}

const TimeLayout = "2006-01-02 15:04:05"

var (
	Local = time.FixedZone("CST", 8*3600)
)

func GetNow() string {
	now := time.Now().In(Local).Format(TimeLayout)
	return now
}
func TestGetNow(t *testing.T) {
	fmt.Println(GetNow())
}

func TimeFormat(s string) string {
	result, err := time.ParseInLocation(TimeLayout, s, time.Local)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	return result.In(Local).Format(TimeLayout)

}
func TestTimeFormat(t *testing.T) {
	TimeFormat("2006-01-02 15:04:05") //转化所需模板2006-01-02 15:04:05为固定值，调换其它都会出错
}

func Md5(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}
func TestMd5(t *testing.T) {
	passwd := Md5("123456")
	fmt.Println(passwd)
}
