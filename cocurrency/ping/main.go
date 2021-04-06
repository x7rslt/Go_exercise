package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

//fake cocurrency
func main() {
	for {
		fmt.Println("Ping")
		ping()
	}
}

var tokens = make(chan struct{}, 200000000)

func ping() {
	tokens <- struct{}{}
	cmd := exec.Command("ping", "47.97.69.185")
	if err := cmd.Run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	<-tokens
	fmt.Println(time.Now().Clock())
}
