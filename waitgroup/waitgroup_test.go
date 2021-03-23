package waitgroup_test

import (
	"fmt"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func GenerateString(n int, ch chan<- string) {
	defer close(ch)

	for i := 0; i < n; i++ {
		ch <- fmt.Sprintf("String %d", i)
	}
}

func PrintString(s string) {
	defer wg.Done()
	fmt.Printf("Content is %s\n", s)
}

func TestWaitgroup(t *testing.T) {
	ch := make(chan string)

	go GenerateString(10, ch)

	for {
		select {
		case s, ok := <-ch:
			if !ok {
				wg.Wait()
				fmt.Println("Processing finished.")
				return
			}
			wg.Add(1)
			go PrintString(s)
		}
	}
}
