package thread_test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

var counter int = 0

func Count(lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Println(counter)
	lock.Unlock()

}

//一个低效的协程方式，仅作演示
func TestLowConcurrency(t *testing.T) {
	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go Count(lock)
	}
	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}

}
