package mylog_test

import (
	"fmt"
	"os"
	"testing"
)

type FuncForTest func(i int)

func TestLog(t *testing.T) {
	_, err := os.Stat("./")
	fmt.Println(err, os.ModePerm)
	t := FuncForTest{func(i int) { fmt.Println(i) }}
	t
}
