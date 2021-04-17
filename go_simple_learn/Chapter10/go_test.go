package test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func Food(c chan<- string) {
	c <- "noodle"
}

//Select 用法
func TestChannel(t *testing.T) {
	c := make(chan string)
	go Food(c)
	select {
	case content := <-c:
		fmt.Println(content)
	}
}

//Channel close

func FoodMenu(c chan string) {
	foods := []string{"meat", "fish", "crab", "soup"}
	for _, i := range foods {
		c <- i
	}
	close(c)
}
func TestCloseChannel(t *testing.T) {
	c := make(chan string)
	go FoodMenu(c)
	for i := range c {
		fmt.Println(i)
	}
}

//竞态
var sum int

func Sum(i int) {
	sum += i
}
func GetSum() int {
	return sum
}
func TestCompete(t *testing.T) {
	for i := 0; i < 10; i++ {
		go Sum(i) //不同状态的goroutine引用同一个数据，容易重复引用，导致结果错误
	}
	fmt.Println(GetSum())
}

//避免竞态
var sum2 int
var m sync.Mutex

func Sum2(i int) {
	m.Lock()
	defer m.Unlock()
	sum2 += i
}

func GetSum2() int {
	m.Lock()
	defer m.Unlock()
	return sum2
}

func TestCompeteSolve(t *testing.T) {
	for i := 0; i < 10; i++ {
		go Sum2(i)
	}
	runtime.Gosched()
	fmt.Println(GetSum2())
}
