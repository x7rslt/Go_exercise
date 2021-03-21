package switch_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSwitch(t *testing.T) {
	rand.Seed(time.Now().UnixNano()) //生成随机数
	switch num := rand.Intn(30); num {
	case 1:
		fmt.Println("this is 1")
	case 2:
		fmt.Println("this is 2")
	default:
		fmt.Println("this is no you choice.")
	}
}

func TestSwitch2(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	switch num := rand.Intn(100); {
	case num < 10:
		fmt.Println(1)
	case num == 11:
		fmt.Println(11)
	default:
		fmt.Println("no rand number as your wish.")
	}
}

func TestSwitch3(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	num := rand.Intn(100)
	switch {
	case num < 10:
		fmt.Println("num is low")
	case num > 50:
		fmt.Println("num is big")
	default:
		fmt.Println("num is good.")
	}
}
