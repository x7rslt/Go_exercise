package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Singleton makes sure there is only one "single" instance across the system.
// by binding to a tcp resource as specified by addr, eg. "127.0.0.1:51337".
//Preventing duplicate process 防止同个应用多次运行
//方法不好用，要通过连接访问才知道占用短空，fmt.Print不输出，没搞懂为啥？
func Singleton(addr string) {
	go singletonServer(addr)
	for {
		// wait and confirm that server was started successfully
		pid, err := getSingletonPID(addr)
		if err == nil && pid == strconv.Itoa(os.Getpid()) {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func singletonServer(addr string) {
	fmt.Println("Current PID:", os.Getpid())
	// Listen for incoming connections.
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, strconv.Itoa(os.Getpid()))
	})
	err := http.ListenAndServe(addr, nil)
	fmt.Println("SingletonServer err:", err) //为什么没有输出显示？
	if err != nil {
		pid, err := getSingletonPID(addr)
		if err != nil {
			log.Fatalln("agent already running, error on retriving pid", err)
		}

		log.Fatalln("agent already running on pid", pid)
	}
}

func getSingletonPID(addr string) (string, error) {
	resp, err := http.Get("http://" + addr + "/")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	fmt.Println("Page content:", string(body))
	fmt.Println("PID :", os.Getpid())
	return string(body), nil
}

func main() {
	singletonServer("127.0.0.1:51337")

}
