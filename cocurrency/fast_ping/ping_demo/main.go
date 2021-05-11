package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/tatsushid/go-fastping"
)

//依旧未解决ping的好方法
func main() {
	for i := 1; i < 255; i++ {
		ip := "120.55.49." + strconv.Itoa(i)
		Ping(ip)
	}

}

func Ping(ip string) {
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", ip)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
		fmt.Println(p.OnRecv)
	}
	p.OnIdle = func() {
		fmt.Println("finish")
	}
	err = p.Run()
	if err != nil {
		fmt.Println(err)
	}
}
