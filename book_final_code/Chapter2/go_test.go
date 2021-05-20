package test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	name := "Peter"
	fmt.Println(strings.ToUpper(name))
}

//3中常用的字符拼接方法
func TestStringCombine(t *testing.T) {
	name := "Peter"
	phone := "13000000000"
	//1
	user1 := strings.Join([]string{name, phone}, ":")

	//2
	user2 := fmt.Sprintf("%s:%s", name, phone)

	//3
	user3 := bytes.Buffer{}
	user3.WriteString(name)
	user3.WriteString(":")
	user3.WriteString(phone)

	fmt.Println(user1, user2, user3.String())
}

//删除空格、替换值

func TestTrim(t *testing.T) {
	foods := "  Rice,nodle,fish"
	foods2 := strings.Trim(foods, " ")
	foods3 := strings.Replace(foods, "fish", "beef", 1)
	fmt.Println(foods, foods2, foods3)
}
