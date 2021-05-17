package test

import (
	"fmt"
	"sync"
	"testing"
)

func Do(wg *sync.WaitGroup, number *int) {
	defer mu.Unlock()
	mu.Lock()
	*number++
	wg.Done()
}

var (
	wg     sync.WaitGroup
	mu     sync.Mutex
	number int
)

func TestLock(t *testing.T) {
	fmt.Println("Start\n-------------")
	wg.Add(10000000)
	for i := 0; i < 10000000; i++ {
		go Do(&wg, &number)
	}
	wg.Wait()
	fmt.Println("Count result:", number)

}
