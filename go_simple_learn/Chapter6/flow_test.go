package test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCompare(t *testing.T) {
	a := 1
	b := 2
	if a > 0 && b < 3 {
		fmt.Println("Result true.")
	} else {
		fmt.Println("Result false.")
	}

	if a > 4 || b < 4 {
		fmt.Println("Result2 true, too.")
	}
}

func TestSwitch(t *testing.T) {
	a := 0
	switch a {
	case 0:
		fmt.Println(reflect.TypeOf(a))
	}
}

func PrintType(t interface{}) {
	switch t.(type) {
	case int:
		fmt.Println("Int")
	case string:
		fmt.Println("String")
	default:
		fmt.Println("Point")
	}
}
func TestSwitchType(t *testing.T) {
	PrintType(1)
	PrintType("hello")
}

func TestSwitchOr(t *testing.T) {
	i := 10
	switch i {
	case 1, 10:
		fmt.Println("i is 0 or 10")
	}
}
