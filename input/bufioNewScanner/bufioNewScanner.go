package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("Please input what you what:")
	for input.Scan() {
		if len(input.Text()) != 0 { //判断无输入时，退出
			fmt.Println("You input:", input.Text())
		} else {
			break
		}
	}
}
