package test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
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
	sum = sum + i
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
	//fmt.Println(sum2)
}

func GetSum2() int {
	m.Lock()
	defer m.Unlock()
	return sum2
}

func TestCompeteSolve(t *testing.T) {
	for i := 0; i < 10; i++ {
		go Sum2(i)
		runtime.Gosched()
	}
	fmt.Println("Sum:", GetSum2())
}

//waitgroup

var wg sync.WaitGroup

func TestWg(t *testing.T) {
	fmt.Println("可以点菜了吗？")

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("老王还没到")
		time.Sleep(1 * time.Second)
		fmt.Println("老王到了")
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("老张还没到")
		time.Sleep(1 * time.Second)
		fmt.Println("老张到了")
	}()
	fmt.Println("等都到齐了")
	wg.Wait()
	fmt.Println("所有人都到了，可以点菜了！")

}
