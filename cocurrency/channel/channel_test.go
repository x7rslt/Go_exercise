package channel_test

import (
	"fmt"
	"testing"
	"time"
)

func TestChannelTimeout(t *testing.T) {
	ch := make(chan int)
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1e9)
		timeout <- true
	}()

	select {
	case <-ch:
	case <-timeout:
		fmt.Println("Time out.")
	}
}

var timeout chan bool

func timer() {
	timeout = make(chan bool, 1)
	go func() {
		time.Sleep(1e9)
		timeout <- true
	}()
}
func TestChannelTimeout2(t *testing.T) {
	ch := make(chan int)

	timer()
	select {
	case <-ch:
	case <-timeout:
		fmt.Println("Time out.")
	}
}
