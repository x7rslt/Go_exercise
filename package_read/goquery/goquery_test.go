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
