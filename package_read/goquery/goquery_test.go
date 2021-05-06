package test

import (
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func WebScrap() {
	res, err := http.Get("http://www.baidu.com")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatal("Response err.")
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	title := doc.Find("title").Text()
	fmt.Println(title)
}
func TestGoquery(t *testing.T) {
	WebScrap()
}

func TestFunc(t *testing.T) {
	add := func(x, y int) {
		fmt.Println(x + y)

	}
	add(1, 2)

	add2 := func(x, y int) int {
		return x + y

	}(2, 3)
	fmt.Println(add2)
}
