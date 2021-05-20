package test

import (
	"fmt"
	"testing"
)

func FamousPlace(places *[]string) {
	for i, j := range *places {
		fmt.Println(i, j)
	}
	(*places)[0] = "anhui"
	fmt.Println(places)
}

func TestArrayPoint(t *testing.T) {
	places := &[]string{"hangzhou", "shanghai", "beijing"}
	FamousPlace(places)

}

//函数变量
type Hi func(num string)

func Hello(num string) {
	fmt.Println(num, "位客人，欢迎光临")
}
func HelloDongbei(num string) {
	fmt.Println(num, "位兄弟，欢迎光临")
}

func SayHello(num string, hi Hi) {
	hi(num)
}

func TestFuncVar(t *testing.T) {
	customs := "东北人"
	num := "3"
	if customs == "东北人" {
		SayHello(num, HelloDongbei)
	} else {
		SayHello(num, Hello)
	}
}
