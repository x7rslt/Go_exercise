package channel_test

import (
	"fmt"
	"testing"
)

func Counter(out chan<- int) {
	for i := 0; i <= 10; i++ {
		out <- i
	}
	close(out) //close只能用来关闭发送方的通道，试图关闭只能接受的通道会报错。
}

func Square(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func Printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
func TestDirection(t *testing.T) {
	natures := make(chan int)
	squares := make(chan int)

	go Counter(natures)
	go Square(squares, natures)
	Printer(squares)

}
