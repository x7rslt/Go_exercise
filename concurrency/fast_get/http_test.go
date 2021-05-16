package test

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"testing"
)

func TestVisit(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(10)
	f, err := os.Open("./domain")
	if err != nil {
		log.Println("Targets file error:", err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		go Visit("http://"+scanner.Text(), &wg)
	}
	wg.Wait()

}

func Visit(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Visit web error:", err)
	} else {
		fmt.Println(resp.StatusCode)
	}
	wg.Done()
}
