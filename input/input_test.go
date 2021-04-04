package input_test

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

//To do :如何进行交互式测试
func TestBufio(t *testing.T) {
	input := bufio.NewScanner(os.Stdin) //注意：Testing模式没有console交互，所以这里input功能不起作用。
	for input.Scan() {
		fmt.Println(input.Text())
	}
}

func TestFmtScanf(t *testing.T) {
	var input string

	fmt.Println("Input something:")
	fmt.Scanf("%s", &input) //注意：Testing模式没有console交互，所以这里input功能不起作用。
	fmt.Println("Your input :", input)
}
