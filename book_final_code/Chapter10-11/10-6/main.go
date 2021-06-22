package main

import (
	"fmt"
	"log"
)

func First(ch chan<- bool) {
	ch <- true
	fmt.Println("买好菜")
	close(ch)
}

func Second(ch1 <-chan bool, ch2 chan<- bool) {
	<-ch1
	ch2 <- true
	fmt.Println("做菜")
	close(ch2)
}

func Eat(ch <-chan bool) {
	for i := range ch {
		log.Println(i)
		fmt.Println("开饭啦：")
	}
}

func main() {
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	go First(ch1)
	go Second(ch1, ch2)
	Eat(ch2)
	fmt.Println("吃好啦！") //注意顺序
	fmt.Println("刷锅去吧...")
}
