package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

func Scan(i int, wg *sync.WaitGroup) {
	port := strconv.Itoa(i)
	host := "218.75.39.42:" + port
	_, err := net.DialTimeout("tcp", host, time.Second*5)
	if err != nil {
		//log.Println("Port Close:", port)
	} else {
		fmt.Println("Open:", port)
	}
	wg.Done()
}

func main() {
	fmt.Println("Scan start:")
	t1 := time.Now()
	var wg sync.WaitGroup
	wg.Add(65536)
	for i := 0; i <= 65535; i++ {
		go Scan(i, &wg)
	}
	wg.Wait()
	fmt.Println("Port Scan time:", time.Since(t1))
}
