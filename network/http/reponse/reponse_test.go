package response_test

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"
)

func TestResGet(t *testing.T) {
	res, err := http.Get("https://www.baidu.com")
	if err != nil {
		log.Fatal(err)
	}
	content, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Status, len(string(content)))
	fmt.Printf("Status %s,Proto %s", res.Status, res.Proto)
	fmt.Printf("Header -data %q\n", res.Header["Date"])
	fmt.Println("Res Body:", res.Body)
	fmt.Println("ContentLength", res.ContentLength)
	fmt.Println("Request:", res.Request, res.Request.Method)
	fmt.Println("Tls:", res.TLS.ServerName)

	fmt.Println(string(content))
}
