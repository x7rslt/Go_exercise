package test

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"testing"
	"time"
)

func Scan(i int, wg *sync.WaitGroup, ports chan string) {
	//defer close(ports)
	port := strconv.Itoa(i)
	host := "218.75.39.42:" + port
	_, err := net.DialTimeout("tcp", host, time.Second*1)
	if err != nil {
		ports <- "close"
	} else {
		fmt.Println("Open:", port)
		ports <- port
	}
	wg.Done()
}

func TestScan(t *testing.T) {
	portlist := []string{}
	ports := make(chan string)
	fmt.Println("Scan start:")
	t1 := time.Now()
	var wg sync.WaitGroup
	wg.Add(65536)
	for i := 0; i <= 65535; i++ {
		go Scan(i, &wg, ports)
		go func() {
			portre := <-ports
			if portre != "close" {
				portlist = append(portlist, portre)
			}
		}()

	}
	//扫描时间变长，待优化

	wg.Wait()
	fmt.Println("Port Scan time:", time.Since(t1))
	fmt.Println("Portlist:", portlist)
}
