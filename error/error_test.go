package error_test

import (
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	for i := 10; i > -10; i-- {
		if i == 0 {
			fmt.Errorf("Number can't is %d", i)
			continue //必须加contine，跳过i==0的情况，不然依旧报错
		}
		fmt.Println(100 / i)

	}
}
