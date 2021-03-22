package interface_test

import (
	"fmt"
	"testing"
)

type People interface {
	Resume()
}

type Man struct {
	sex string
	age int
}

func (p Man) Resume() {
	fmt.Printf("This people is %s,%d years old.", p.sex, p.age)
}

func TestInterface(t *testing.T) {
	var xiao Man = Man{"Femal", 27}
	xiao.Resume()
}
