package point_test

import (
	"fmt"
	"testing"
)

type People struct {
	name string
	age  int
}

func (p *People) Hello() {
	fmt.Println("Hi,I'm ", p.name, p.age, " years old.")
}

func TestPoint(t *testing.T) {
	var xiao People
	fmt.Println(xiao)
	var shuai *People
	fmt.Println(shuai)
	xiao.Hello()
	//shuai = &People{"xiao", 12}
	shuai.Hello()
}
