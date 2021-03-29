package regex_test

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"testing"
)

//添加cookie
func TestCookie(t *testing.T) {
	url := "https://www.qixintong.cn/company/68394f714f6f4f4e4b4d43?utm_source=down_fav"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	cookie1 := &http.Cookie{Name: "_ga", Value: "GA1.2.1598405589.1603247431"}
	req.AddCookie(cookie1)
	cookie2 := &http.Cookie{Name: "utm_source", Value: "SEM"}
	req.AddCookie(cookie2)
	cookie3 := &http.Cookie{Name: "_gid", Value: "GA1.2.1626449279.1616988487"}
	req.AddCookie(cookie3)
	cookie4 := &http.Cookie{Name: "Hm_lvt_8ade5034b029fd3234640e22eb8ac190", Value: "1616988487"}
	req.AddCookie(cookie4)
	cookie5 := &http.Cookie{Name: "sessionid", Value: "k139mx5tw730ie9ohf6zlgietf1jwoyh"}
	req.AddCookie(cookie5)
	cookie6 := &http.Cookie{Name: "utype", Value: "1"}
	req.AddCookie(cookie6)
	cookie7 := &http.Cookie{Name: "tel", Value: "155****1522"}
	req.AddCookie(cookie7)
	cookie8 := &http.Cookie{Name: "Hm_lpvt_8ade5034b029fd3234640e22eb8ac190", Value: "1617007167"}
	req.AddCookie(cookie8)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	fmt.Println(req.Cookies())
	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))

}
func Fetch(u string) (string, error) {
	fmt.Println(u)
	res, err := http.Get(u)
	if err != nil {
		return "", err
	}
	if res.StatusCode != http.StatusOK {
		fmt.Printf("Connect error:%d", res.StatusCode)
	}
	content, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return "", err
	}
	result := string(content)
	re := regexp.MustCompile(`<a rel="nofollow" target="_blank" href="(.*)">`)
	url := re.FindString(result)
	fmt.Println(result)
	fmt.Println(url)
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
