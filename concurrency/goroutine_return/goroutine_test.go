package test

import (
	"fmt"
	"testing"
)

//自定义类型，通过通道统一传递值
type Result struct {
	name   string
	length int
}

func PersonCount(userlist []string, ch chan Result) {
	defer close(ch)
	result := new(Result)
	for _, i := range userlist {
		result.name = i
		result.length = len(i)
		ch <- *result
	}

}

func TestChannelReturn(t *testing.T) {
	userlist := []string{
		"lehehe", "muhahah", "kaixin",
	}
	ch := make(chan Result)
	go PersonCount(userlist, ch)
	resultlist := []Result{}
	for i := range ch {
		resultlist = append(resultlist, i)
	}
	fmt.Println(resultlist)
}
