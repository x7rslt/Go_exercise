package main

import (
	"flag"
	"fmt"
)

var (
	ip      string
	port    int
	verbose bool
)

func init() {
	flag.StringVar(&ip, "IP", "127.0.0.1", "target IP")
	flag.IntVar(&port, "Port", 80, "Port")
	flag.BoolVar(&verbose, "Verbose", true, "verbosity")
}

func main() {
	flag.Parse()
	fmt.Printf("Scan %s:%d\n", ip, port)
	if verbose {
		fmt.Println("Attack!")
	}
}
