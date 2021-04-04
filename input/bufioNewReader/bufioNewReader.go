package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter something:")
	value, _ := reader.ReadString('\n')
	fmt.Println("Your input:", value)
}
