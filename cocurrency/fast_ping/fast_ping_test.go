package test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/go-ping/ping"
)

//待解决Error ：ping 未存活状态的ip陷入死循环
func TestFastPint(t *testing.T) {
	i := 1

	fmt.Println(strconv.Itoa(i))
}

func Ping() {

	pinger, err := ping.NewPinger("120.55.49.8")
	if err != nil {
		panic(err)
	}
	pinger.Count = 3
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		panic(err)
	}
	stats := pinger.Statistics() // get send/receive/duplicate/rtt stats
	fmt.Println(stats)
}
