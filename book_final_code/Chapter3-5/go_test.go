package test

import (
	"fmt"
	"testing"
)

//常见的slice错误，指针方式重复赋值
func TestSliceError(t *testing.T) {
	foods := []string{"北京", "上海", "杭州", "广东", "深圳"}
	foods2 := make([]*string, len(foods))
	for i, value := range foods {
		foods2[i] = &value
	}
	fmt.Println(foods, []string{*foods2[0], *foods2[1]})
}

//修改：
func TestSlice(t *testing.T) {
	foods := []string{"北京", "上海", "杭州", "广东", "深圳"}
	foods2 := make([]*string, len(foods))
	for i, _ := range foods {
		foods2[i] = &foods[i]
	}
	fmt.Println(foods, []string{*foods2[0], *foods2[1]})
}
