package runtimeGosched_test

import (
	"fmt"
	"runtime"
	"testing"
)

func Haha() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hahaha")
	}
}

func TestGosched(t *testing.T) {
	go Haha()
	runtime.Gosched() //作用是让当前goroutine让出CPU，好让其它的goroutine获得执行的机会。同时，当前的goroutine也会在未来的某个时间点继续运行。
	fmt.Println("Done")
}
