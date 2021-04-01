package concurrencyz

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func Crawl(url string) {
	req, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(req.StatusCode, string(content)[:100])

}

func TestCon(t *testing.T) {
	url := "http://www.baidu.com"
	done := make(chan bool)
	for {
		go func() {
			Crawl(url)
		}()
		done <- true
	}
}
