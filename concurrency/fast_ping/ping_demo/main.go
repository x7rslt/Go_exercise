package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

//依旧未解决ping的好方法
func main() {
	aliveIP := []string{}
	deadIP := []string{}

	for i := 3; i < 5; i++ {
		ip := "120.55.49." + strconv.Itoa(i)
		if CheckIp(ip) {
			aliveIP = append(aliveIP, ip)
		} else {
			deadIP = append(deadIP, ip)
		}
	}
	fmt.Println("Alive ip number :", len(aliveIP))
	fmt.Println("Dead ip number :", len(deadIP))

}

func CheckIp(ip string) bool {
	out, _ := exec.Command("ping", ip, "-c 3").Output()
	if strings.Contains(string(out), "Request timeout for icmp_seq") {
		return false
	} else {
		return true
	}
}
