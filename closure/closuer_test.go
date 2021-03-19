package go_example_test

import (
	"fmt"
	"testing"
)

//闭包
func Add(m int) func(n int) int {
	return func(n int) int {
		num := n * m //运算要和后面的return拆分
		return num
	}
}

func TestClosures(t *testing.T) {
	num := Add(10)
	fmt.Println(num(10))

}


