package args_test

import (
	"fmt"
	"testing"
)

func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg)
		case string:
			fmt.Println(arg)
		default:
			fmt.Println("blank")
		}
	}
}
func TestArgs(t *testing.T) {
	a := 1
	b := "tset"
	c := ""
	MyPrintf(a, b, c)

}
