package channel_test

import (
	"fmt"
	"runtime"
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
func TestChannelTimeoutFunc(t *testing.T) {
	ch := make(chan int)

	timer()
	select {
	case <-ch:
	case <-timeout:
		fmt.Println("Time out.")
	}
}

//问题：接收的数据只有1、3、5、7、9？为什么？
func TestChannel(t *testing.T) {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		go func(v int) {
			//fmt.Println(i)
			ch <- v
		}(i)
	}

	go func() {
		for range ch {
			fmt.Println(<-ch)
		}
	}()

	runtime.Gosched()
	fmt.Println("Done")
}

//问题：输出莫名其妙的数字？
func TestChannel2(t *testing.T) {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		//fmt.Println(i)
		go func(v int) {
			//fmt.Println(i)
			ch <- v
		}(i)
	}

	for range ch {
		timer()
		select {
		case <-ch:
			fmt.Println(<-ch)
		case <-timeout:
			fmt.Println("Timeout.")
			fmt.Println("Done")
		}

	}

	runtime.Gosched()
	fmt.Println("Done")
}
