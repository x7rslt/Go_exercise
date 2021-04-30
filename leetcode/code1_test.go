package code1_test

import (
	"fmt"
	"sort"
	"testing"
)

func Test1(t *testing.T) {
	var nums = []int{2, 11, 7}
	var targets = 9

	for j := 0; j < len(nums); j++ {
		for i := j; i < len(nums); i++ {
			if nums[j]+nums[i] == targets {
				fmt.Println(j, i)
			}

		}

	}

}

//错题集：不知道为什么结果正负错误
/*
解析：
因为取值范围：
int8: -128 ~ 127
int16: -32768 ~ 32767
int32: -2147483648 ~ 2147483647
int64: -9223372036854775808 ~ 9223372036854775807
*/
func arraySign(nums []int) int {
	product := 1
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
		product *= nums[i]
	}
	fmt.Println(product)
	if product > 0 {
		return 1
	} else if product < 0 {
		return -1
	} else {
		return 0
	}
}

func Test2(t *testing.T) {
	nums := []int{9, 72, 34, 29, -49, -22, -77, -17, -66, -75, -44, -30, -24}
	fmt.Println(arraySign(nums))
}

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	result := 0
	for i := 0; i < len(costs); i++ {
		if coins >= costs[i] {
			coins -= costs[i]
			fmt.Println(costs[i])
			result++
			continue
		}
		break
	}
	return result
}
func Test3(t *testing.T) {
	consts := []int{1, 3, 2, 4, 1}
	coins := 12
	fmt.Println(maxIceCream(consts, coins))
}
