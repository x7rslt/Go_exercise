package test

import (
	"fmt"
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
