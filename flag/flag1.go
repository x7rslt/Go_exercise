package main

import (
	"flag"
	"fmt"
)

func main() {

	// Declare flags
	// flag string
	ipstr := flag.String("ip", "127.0.0.1", "target IP")

	//flag int
	var port int
	flag.IntVar(&port, "port", 8080, "Port")

	//flag bool
	ipbool := flag.Bool("verbose", true, "verbosity")

	// Parse flags
	flag.Parse()

	// Flag print:Hacking ip
	fmt.Printf("Hacking %s:%d!\n", *ipstr, port)

	// Display progress if verbose flag is set
	if *ipbool {
		fmt.Printf("Pew pew!\n")
	}
}
