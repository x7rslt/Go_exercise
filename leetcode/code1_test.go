package code1_test

import (
	"fmt"
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
