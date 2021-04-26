package test

import (
	"fmt"
	"testing"
)

//空接口可以保存任何值
type i interface{}

func Say(s i) {
	fmt.Printf("%v,%T\n", s, s)
}

func TestNilInterface(t *testing.T) {
	var j interface{}
	Say(j)

	a := 1
	Say(a)

	b := "Hi"
	Say(b)
}

//使用指针接收者实现接口
//对于方法接收者，每个具体类型的方法接收者就是这个类型本身。当调用方法时，每次调用都要有一个接收者，如果频繁调用，
//那么就会有多个接收者，这会极大的浪费内存，毕竟计算机的内存是有限的。如何解决这个问题？
//建议使用（*）取地址的操作，就是每次都去这个接收的地址那里找接收者，这样无论调用多少次方法，接收者都是同一个地址，
//即同一个对象，这样可以节省计算机的内存。
type Studen interface {
	Study() bool
	Subject(lesson string) bool
}

type HighSchoolStuden struct {
	name string
	age  int
}

func (h *HighSchoolStuden) Study() bool {
	fmt.Println(h.name, "study everday.")
	fmt.Printf("%p\n", h)
	return true
}

func (h *HighSchoolStuden) Subject(lesson string) bool {
	fmt.Printf("%s is %d years old,favorite like subject is %s.\n", h.name, h.age, lesson)
	fmt.Printf("%p\n", h)
	return true
}

func GoodStudent(h *HighSchoolStuden, lesson string) {
	h.Study()
	h.Subject(lesson)
}

func TestPointerInterface(t *testing.T) {
	s := &HighSchoolStuden{
		name: "Xiao",
		age:  28,
	}
	GoodStudent(s, "Computer Science")
}
