package main

import (
	"flag"
	"fmt"
)

func main() {
	// Set flag
	_ = flag.Int("flag1", 0, "flag1 description")
	// Parse all flags
	flag.Parse()
	// 列出命令参数的两种方式
	//方式一：
	fmt.Println("Method 1:")
	for _, v := range flag.Args() {
		fmt.Println(v)
	}

	fmt.Println("Method 2:")
	// 方式二：
	for i := 0; i < flag.NArg(); i++ {
		fmt.Println(flag.Arg(i))
	}
}
