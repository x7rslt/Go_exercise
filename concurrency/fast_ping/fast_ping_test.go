package test

import (
	"fmt"
<<<<<<< HEAD
	"net"
	"os"
=======
	"strconv"
>>>>>>> 0237545dcafcb5112f488bbbf665beec237a545c
	"testing"
	"time"

	"github.com/tatsushid/go-fastping"
)

//待解决Error ：ping 未存活状态的ip陷入死循环
func TestFastPint(t *testing.T) {
	i := 1

	fmt.Println(strconv.Itoa(i))
}

func Ping() {
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
	}
	p.OnIdle = func() {
		fmt.Println("finish")
	}
	err = p.Run()
	if err != nil {
		fmt.Println(err)
	}
}
