package download_test

import (
	"io"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestDownload(t *testing.T) {
	website, err := os.Create("download")
	if err != nil {
		log.Fatal(err)
	}
	defer website.Close()

	url := "http://www.baidu.com"
	response, err := http.Get(url)
	defer response.Body.Close()

	numBytesWrite, err := io.Copy(website, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Download %d byte file.\n", numBytesWrite)
}
