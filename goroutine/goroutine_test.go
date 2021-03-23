package goroutine_test

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	go func() {
		fmt.Println("Hello Goroutine.")
	}()
	time.Sleep(time.Second) //Delay main() finish
	fmt.Println("Main finished.")
}

func TestGoroutineDone(t *testing.T) {
	fmt.Println("Start a groutine.")
	done := make(chan struct{}) //定义一个struct{}通道来同步goroutine状态
	go func() {
		fmt.Println("This is a groutine.")
		done <- struct{}{} //发送信号说明goroutine运行结束
	}()
	<-done //结束goroutine的信号
	fmt.Println("Goroutine done.")

}
