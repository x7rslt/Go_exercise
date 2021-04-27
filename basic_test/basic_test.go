package basic_test

import (
	"fmt"
	"log"
	"testing"
)

//面向对象
type IntCompare int

func (a IntCompare) Big(b IntCompare) bool {
	return a > b
}
func TestIntC1(t *testing.T) {
	a := IntCompare(1)
	b := IntCompare(2)
	result := a.Big(b)
	fmt.Println(result)
}

//面向过程
func Big(a, b int) bool {
	return a > b
}
func TestIncC2(t *testing.T) {
	a, b := 2, 1
	result := Big(a, b)
	fmt.Println(result)
}

//错误的定义a值增加的方法
type Integer int

func (a Integer) Add(b Integer) {
	a += b
}

func TestIntAdd(t *testing.T) {
	a := Integer(1)
	b := Integer(2)
	a.Add(b)
	fmt.Println(a)
}

//正确定义a值增加的方法，要用指针

func (a *Integer) Add2(b Integer) {
	*a += b
}
func TestIntAdd2(t *testing.T) {
	a := Integer(1)
	b := Integer(2)
	a.Add2(b)
	fmt.Println(a)
}

var str string
var str2 = ""

func TestVar(t *testing.T) {
	str3 := "" //注：只能定义的局部变量
	fmt.Println(str, str2, str3)
}

type Person struct {
	name string
}

func TestPointer(t *testing.T) {
	p := Person{"lehehe"}
	po := &p
	fmt.Println(po, *po) //注：&是指针，*是变量值

	p2 := &Person{"hahaha"}
	fmt.Println(p2.name) //指针成员变量可直接引用，编译器会自动隐形p转为（*p2）
	fmt.Println((*p2).name)
}

//[primary] 关于接口和类的说法，下面说法正确的是（）
//A. 一个类只需要实现了接口要求的所有函数，我们就说这个类实现了该接口
//B. 实现类的时候，只需要关心自己应该提供哪些方法，不用再纠结接口需要拆得多细才合理
//C. 类实现接口时，需要导入接口所在的包   //测试：将interface独立分包，不需要导入即可引用
//D. 接口由使用方按自身需求来定义，使用方无需关心是否有其他模块定义过类似的接口
//
//参考答案：ABD

//这里说的类就是Go的结构体
type le interface {
	ha() string
}
type kaixin struct {
	name string
}

func (l *kaixin) ha(s string) string {
	return fmt.Sprintf("%s %s.", l.name, s)
}

func TestInterface(t *testing.T) {
	name := kaixin{"xiao"}
	status := name.ha("happy.")
	fmt.Println(status)
}

func TestStr(t *testing.T) {
	str := "123" + `abc`
	fmt.Println(str)
}

func TestFor(t *testing.T) {
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}
	for i := 0; i < 100; i++ {
		fmt.Println("lehehe")
	}
}

func TestQuote(t *testing.T) {
	a := "\n"
	b := 'd'
	fmt.Printf("Type : a %T,b %T", a, b)
}

//值引用
func TestValue(t *testing.T) {
	a := [3]int{1, 2, 3}
	b := a
	b[1]++
	fmt.Print(a, b)
}

func TestValue2(t *testing.T) {
	a := [3]int{1, 2, 3}
	b := &a
	b[1]++
	fmt.Println(a, *b)
}

func TestNil(t *testing.T) {
	var both_nil interface{} = nil
	fmt.Println("both_nil == nil", both_nil == nil)

	var nil_ptr *string
	var val_nil interface{} = nil_ptr

	fmt.Println("val_nil==nil", val_nil == nil)
}

//指针
func TestPointer2(t *testing.T) {
	var x int
	p := &x
	fmt.Println(x, *p, p)
}

func TestLogFatal(t *testing.T) {
	a := 1
	if a == 2 {
		log.Fatal("a = ", a)
	}
	log.Println("a 不等于2")
}
