package main

import (
	"flag"
	"fmt"
	"os"
)

//Execute:go run flag_subcommand.go add -x 23 -y 45
func main() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	x := addCmd.Int("x", 1, "x")
	y := addCmd.Int("y", 2, "y")

	sqrtCmd := flag.NewFlagSet("sqrt", flag.ExitOnError)
	z := sqrtCmd.Int("z", 7, "z")

	if len(os.Args) < 2 {
		fmt.Println("expected 'add' or 'sqrt' subcommands")
		os.Exit(0)
	}
	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		//fmt.Println(x, y) 打印指针地址
		fmt.Println(*x, *y)
	case "sqrt":
		sqrtCmd.Parse(os.Args[2:])
		fmt.Println(*z)
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}
