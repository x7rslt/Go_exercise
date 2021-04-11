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
