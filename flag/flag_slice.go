package main

import (
	"flag"
	"fmt"
	"strings"
)

var Add string

//Execute:go run flag4.go -Add f1,f2
func main() {
	flag.StringVar(&Add, "Add", "12,12", "-Add num1,num2")
	flag.Parse()
	value := strings.Split(Add, ",")
	fmt.Println(value)
}
