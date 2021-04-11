package channelcpu_test

import (
	"fmt"
	"runtime"
	"testing"
)

func TestChannelCpu(t *testing.T) {
	num_cpu := runtime.NumCPU()
	fmt.Println(runtime.GOMAXPROCS(num_cpu))
}

func TestCpuThread(t *testing.T) {
	runtime.GOMAXPROCS(2)
	for {
		go func() {
			fmt.Println(0)
		}()
		fmt.Println(1)
	}

}
