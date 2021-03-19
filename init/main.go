package main

import "fmt"

func init() { //init定义后自动执行，不需要在main中引用
	fmt.Println("This is init.")
}

func main() {
	fmt.Println("This is init test.")
}
