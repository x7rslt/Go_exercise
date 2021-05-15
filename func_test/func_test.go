package test

import (
	"fmt"
	"testing"
)

type Pinger struct {
	rec func(s string, i int)
}

func Test(t *testing.T) {
	var p Pinger
	p.rec = func(s string, i int) {
		fmt.Println(s, i)
	}
}
