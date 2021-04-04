package ifelse_test

import (
	"fmt"
	"testing"
)

func TestIf(t *testing.T) {
	a := 10
	if a > 0 {
		fmt.Println("a big than 0")
	} else {
		fmt.Println("a small than 0")
	}
}

func TestElif(t *testing.T) {
	a := 10
	if a > 5 {
		fmt.Println("a big than 5")
		if a > 12 {
			fmt.Println("a big than 12")
		} else if a > 10 {
			fmt.Println("a big than 10")
		} else if a < 10 {
			fmt.Println("a big than 10")
		} else {
			fmt.Println("a is 10")
		}
	}
}
