package main

import (
	"flag"
	"fmt"
	"strings"
)

var Add string

func main() {
	flag.StringVar(&Add, "Add", "12,12", "-Add num1,num2")
	flag.Parse()
	value := strings.Split(Add, ",")
	fmt.Println(value)
}
