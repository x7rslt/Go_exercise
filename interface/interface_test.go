package interface_test

import (
	"fmt"
	"testing"
)

type People interface {
	Resume()
}

type Man struct {
	sex string
	age int
}

func (p Man) Resume() {
	fmt.Printf("This people is %s,%d years old.", p.sex, p.age)
}

func TestInterface(t *testing.T) {
	var xiao Man = Man{"Femal", 27}
	xiao.Resume()
}

//interface可以接受任意值
var emptyInterface interface{}

func Test(t *testing.T) {
	emptyInterface = 10
	fmt.Printf("%T\n", emptyInterface)

	emptyInterface = 10.0
	fmt.Printf("%T\n", emptyInterface)

	emptyInterface = 10.00
	fmt.Printf("%T\n", emptyInterface)

	emptyInterface = "Hello world"
	fmt.Printf("%T\n", emptyInterface)

}

func printType(i interface{}) {
	switch val := i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case float64:
		fmt.Println("float64")
	default:
		fmt.Println("Other:", val)
	}
}

func TestType(t *testing.T) {
	var i interface{}
	printType(i)
	printType(10)
	printType("hello")
}
