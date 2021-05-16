package test

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/go-ping/ping"
)

var (
	mu      sync.Mutex
	deadIP  []string
	aliveIP []string
)

func Test(t *testing.T) {
	t1 := time.Now()
	var wg sync.WaitGroup
	wg.Add(255)
	for i := 1; i <= 255; i++ {

		host := "120.55.49." + strconv.Itoa(i)
		go Ping(host, &wg)

	}
	wg.Wait()
	scantime := time.Since(t1)
	fmt.Println("Ping down.")
	fmt.Printf("Scan spend time: %s\n", scantime)
	fmt.Println("Alive ip:", aliveIP)
	fmt.Println("Dead ip:", deadIP)
	fmt.Println("Deadip number:", len(deadIP))

}

func Ping(host string, wg *sync.WaitGroup) {
	timeout := time.Second * 3
	interval := time.Second
	count := -1
	size := 16
	privileged := false
	pinger, err := ping.NewPinger(host)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}

	// listen for ctrl-C signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			pinger.Stop()
		}
	}()

	//pinger.OnRecv = func(pkt *ping.Packet) {
	//	fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v ttl=%v\n",
	//		pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.Ttl)
	//}
	pinger.OnDuplicateRecv = func(pkt *ping.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v ttl=%v (DUP!)\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.Ttl)
	}
	pinger.OnFinish = func(stats *ping.Statistics) {
		//fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
		//fmt.Printf("%d packets transmitted, %d packets received, %d duplicates, %v%% packet loss\n",
		//	stats.PacketsSent, stats.PacketsRecv, stats.PacketsRecvDuplicates, stats.PacketLoss)
		//fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
		//	stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
		mu.Lock()
		if stats.PacketsRecv != 0 {
			//fmt.Printf("%s is Alive.\n", host)
			aliveIP = append(aliveIP, host)
		} else {
			//fmt.Printf("%s is Dead.\n", host)
			deadIP = append(deadIP, host)
		}
		mu.Unlock()
	}

	pinger.Count = count
	pinger.Size = size
	pinger.Interval = interval
	pinger.Timeout = timeout
	pinger.SetPrivileged(privileged)
	//fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	err = pinger.Run()
	if err != nil {
		fmt.Printf("Failed to ping target host: %s", err)
	}
	wg.Done()
}
