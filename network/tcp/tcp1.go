package main

import (
	"fmt"
	"net"
)

func main() {
	t, err := net.Dial("tcp", "baidu.com:80")
	if err != nil {
		panic(err)
	}
	fmt.Println(*t)
}
