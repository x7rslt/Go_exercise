package permissions_test

import (
	"io"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestPermissions(t *testing.T) {
	newFile, err := os.Create("filetest.html")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	url := "http://www.baidu.com"  //注意url地址要正确，否则请求失败报错
	response, err := http.Get(url) //err判断，防止错误终端运行
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	numBytesWritten, err := io.Copy(newFile, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Download %d byte file.\n", numBytesWritten)
}
