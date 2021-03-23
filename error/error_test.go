package error_test

import (
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	for i := 10; i > -10; i-- {
		if i == 0 {
			fmt.Errorf("Number can't is %d", i)
			continue
		}
		fmt.Println(100 / i)

	}
}
