package test

import (
	"fmt"
	"testing"
)

type ErrNum struct {
	Code    int
	Message string
}

var OK = &ErrNum{0, "OK"}

func TestMyErr(t *testing.T) {
	fmt.Println(OK.Code)
}
