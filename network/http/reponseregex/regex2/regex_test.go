package regex_test

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"testing"
	"time"
)

//添加cookie
func Fetch(u string) (string, error) {
	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	req, _ := http.NewRequest("GET", u, nil)

	cookie1 := &http.Cookie{Name: "_ga", Value: "GA1.2.2028119289.1616856075"}
	req.AddCookie(cookie1)
	cookie2 := &http.Cookie{Name: "utm_source", Value: "SEM"}
	req.AddCookie(cookie2)
	cookie3 := &http.Cookie{Name: "_gid", Value: "GA1.2.676504208.1617441411"}
	req.AddCookie(cookie3)
	cookie4 := &http.Cookie{Name: "Hm_lvt_8ade5034b029fd3234640e22eb8ac190", Value: "1616856075"}
	req.AddCookie(cookie4)
	cookie5 := &http.Cookie{Name: "sessionid", Value: "ci3hvnsh5u92puwag47emqhvjuzh5f8y"}
	req.AddCookie(cookie5)
	cookie6 := &http.Cookie{Name: "utype", Value: "1"}
	req.AddCookie(cookie6)
	cookie7 := &http.Cookie{Name: "tel", Value: "15555001522"}
	req.AddCookie(cookie7)
	cookie8 := &http.Cookie{Name: "Hm_lpvt_8ade5034b029fd3234640e22eb8ac190", Value: "1617441425"}
	req.AddCookie(cookie8)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	//fmt.Println(req.Cookies())
	res, err := client.Do(req)

	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Printf("Connect error:%d", res.StatusCode)
	}
	content, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return "", err
	}
	result := string(content)
	re := regexp.MustCompile(`<a rel="nofollow" target="_blank" href="(.*)">`)
	url := re.FindString(result)
	//fmt.Println(result)
	//fmt.Println(url)
	return url, nil

}

func TestReadScan(t *testing.T) {
	wordfile, err := os.Open(`url`)
	if err != nil {
		log.Fatal(err)
	}
	defer wordfile.Close()

	done := make(chan bool)
	scanner := bufio.NewScanner(wordfile)
	go func() {
		for scanner.Scan() {
			pass := scanner.Text()
			url, err := Fetch(pass)
			fmt.Println(url, err)
		}
		done <- true
	}()
	<-done
}
