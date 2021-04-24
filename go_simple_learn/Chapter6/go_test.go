package test

import (
	"fmt"
	"reflect"
	"testing"
)

//flow if and or
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

//flow switch
func TestSwitchOr(t *testing.T) {
	i := 10
	switch i {
	case 1, 10:
		fmt.Println("i is 0 or 10")
	}
}

//map dict

func TestMap(t *testing.T) {
	var m map[string]string = map[string]string{"xiao": "26", "zhang": "27", "wang": "28"}
	name, ok := m["xiao"]
	if ok {
		fmt.Println(name)
	} else {
		fmt.Println("m no this name.")
	}
	//delete key
	var food map[string]int = map[string]int{"炒鸡蛋": 15, "卤猪蹄": 25, "地锅鸡": 68}
	delete(food, "卤猪蹄")
	delete(food, "海鲜")
	fmt.Println(food)

}
func TestMapAppend(t *testing.T) {
	//map iterate append error:same value
	var m map[string]string = map[string]string{"xiao": "26", "zhang": "27", "wang": "28"}
	var p []*string
	for k, v := range m {
		fmt.Println(k, v)
		p = append(p, &v)
		fmt.Println("#P value :")
		for _, i := range p {
			fmt.Println(*i)
		}
	}

	//Correct:
	var m2 map[string]string = map[string]string{"xiao": "26", "zhang": "27", "wang": "28"}
	var p2 []*string
	for k2, _ := range m2 {
		v2 := m2[k2]
		p2 = append(p2, &v2)
		fmt.Println("#P2 value :")
		for _, i2 := range p2 {
			fmt.Println(*i2)
		}
	}
}
