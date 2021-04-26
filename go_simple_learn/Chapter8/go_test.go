package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

//Json是一种特殊格式的字符串，可以传输和存储数据，在开发中，JSON主要负责
//给前端提供数据，而前端和后端交互的数据格式也是以JSON为主。现在JSON已经成为前端和后端开发的通信桥梁。

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//Marshal 编码
func TestJson(t *testing.T) {
	u := User{
		Name: "xiao",
		Age:  28,
	}
	marshal, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(marshal))
	}
}

//Unmarshal解码
func TestJsonUnmarshal(t *testing.T) {
	var u User
	s := `{"name":"xiao","age":28}`
	err := json.Unmarshal([]byte(s), &u)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v", u)
	}
}

//Point receiver method
//两种需要用到指针接收者方法的情况：
//·在调用方法时，需要更新变量
//·类型的成员很多，占用内存很大，这样的开销会让内存使用率迅速增大
type Animal struct {
	name   string
	action string
}

func (a *Animal) Detail(kind string) {
	a.name = "Jeffy"
	fmt.Printf("The %s's name is %s.\n", kind, a.name)
}

func (a *Animal) Do() {
	fmt.Println(a.name, " can ", a.action)
}
func TestPointMethod(t *testing.T) {
	dog := &Animal{
		name:   "Huzi",
		action: "run",
	}
	dog.Detail("dog")
	dog.Do()

}
