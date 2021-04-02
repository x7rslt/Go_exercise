package url_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/asaskevich/govalidator"
)

func TestUrl(t *testing.T) {
	u := "http://test.com-"
	ub := govalidator.IsURL(u)
	fmt.Println(ub)
	for {
		if !ub {
			log.Fatal("Invalid url.")
		}
		resp, err := http.Get(u)
		if err != nil {
			log.Println(err)
		}
		defer resp.Body.Close()
		fmt.Println("status code:", resp.StatusCode)
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Println(err)
		}
		fmt.Println(body)
	}

}
