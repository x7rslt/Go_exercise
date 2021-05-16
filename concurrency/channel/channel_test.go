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
func TestChannel3(t *testing.T) {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		//fmt.Println(i)
		go func(v int) {
			//fmt.Println(i)
			ch <- v
		}(i)
	}
	close(ch)
	for v := range ch {
		fmt.Println(v)
	}

	runtime.Gosched()
	fmt.Println("Done")
}

func MoreFast(i int, done chan bool) {
	fmt.Println(i)
	time.Sleep(2 * time.Second)
	fmt.Println(i, " time done.")
	done <- true
}

func TestActiveThread(t *testing.T) {
	activeThreads := 0
	maxThreads := 10
	done := make(chan bool)

	num := 0
	for i := 0; i < 20; i++ {
		go MoreFast(num, done)
		activeThreads++

		if activeThreads >= maxThreads {
			<-done
			activeThreads -= 1
		}
		num++
	}

	//等待所有go routine完成，在进行下一步输出“Done”；如果没有这段会在部分go routine未结束时输出 Done。
	for activeThreads > 0 {
		<-done
		activeThreads -= 1
	}
	fmt.Println("Done")
}
