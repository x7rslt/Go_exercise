package enum_test

import (
	"fmt"
	"testing"
)

//项目中总共有四个enum：pay_status、pay_type、sell_status、user_enum,这里仅演示pay_status
type pay_status int

const (
	payed pay_status = 1
	unpay pay_status = 0
)

func (p pay_status) String() {
	switch p {
	case payed:
		fmt.Println("product is payed!")
	case unpay:
		fmt.Println("product is unpay!")
	}
}

func TestPayStatusEnum(t *testing.T) {
	var p pay_status = 0
	p.String()
}
