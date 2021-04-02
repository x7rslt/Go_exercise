package concurrencyz

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

var done chan bool

func Crawl(url string) {
	req, err := http.Get(url)
	if err != nil {
		log.Println("Http get error :", err)
	} else {
		content, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(req.StatusCode, len(content))
		fmt.Println(req.Request.URL)
	}

	done <- true
}

func Url(f string) {
	urlfile, err := os.Open(f)
	if err != nil {
		log.Println(err)
	}
	scanner := bufio.NewScanner(urlfile)
	for scanner.Scan() {
		u := scanner.Text()
		go Crawl(u)
		<-done
	}
}
func TestCon(t *testing.T) {
	urlfile := "urls.txt"
	done = make(chan bool)
	start := time.Now()
	Url(urlfile)
	end := time.Since(start)
	fmt.Println("Crawl spend time:", end)
}
