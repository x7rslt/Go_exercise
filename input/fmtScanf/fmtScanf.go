package main

import "fmt"

func main() {
	var input string

	fmt.Println("Input something:")
	fmt.Scanf("%s", &input) //单次输入,多次输入是用for循环
	fmt.Println("Your input :", input)
	var name string
	var age int
	fmt.Println("Input your name and age:")
	fmt.Scanf("%s %d", &name, &age) //单次输入多个值
	fmt.Printf("Hello %s ,%d years old boy.\n", name, age)
}
