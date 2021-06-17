package func_test

import (
	"errors"
	"fmt"
	"testing"
)

func Divide(a, b int) (int, error) {
	if b != 0 {
		c := a / b
		return c, nil
	} else {
		return 0, errors.New("Func run error:b is 0")
	}
}

func TestFunc(t *testing.T) {
	a := 1
	b := 0
	c, error := Divide(a, b)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(c)
	}
}
