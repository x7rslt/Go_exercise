package channel_test

import (
	"fmt"
	"runtime"
	"testing"
)

func TestChannel(t *testing.T) {
	ch := make(chan int)
	done := make(chan struct{})
	i1 := 10
	go func() {
		ch <- i1
		done <- struct{}{}
	}()
	fmt.Println(<-ch)
	<-done
	fmt.Println("Channel Done.")
}

func TestChannelN(t *testing.T) {
	ch := make(chan int, 10)
	done := make(chan struct{})

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		done <- struct{}{}
	}()

	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
	<-done
	fmt.Println("Main finished.")
}

func TestChannelWR(t *testing.T) {
	ch := make(chan int, 2)
	for i := 0; i < 10; i++ {
		select {
		case ch <- i:
			fmt.Println("Sent to channel", i)
		case i2 := <-ch:
			fmt.Println("Receieved from channel", i2)
		default:
			fmt.Println("None action.")
		}
	}
}

func TestChannelClose(t *testing.T) {
	ch := make(chan int)
	go func() {
		ch <- 1
		close(ch)
	}()

	x, ok := <-ch
	if ok {
		fmt.Println("ok:", ok)
		fmt.Println("Channel ch:", x)
	}

	x, ok = <-ch
	if !ok {
		fmt.Println("ok:", ok)
		fmt.Println("Channel close:", x)
	}
	runtime.Gosched()
	fmt.Println("Channel close test done.")
}
