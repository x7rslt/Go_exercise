package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"regexp"
)

func Gethost(i string) string {
	u, err := url.Parse(i)
	if err != nil {
		log.Fatal(err)
	}
	return u.Host

}
func main() {
	r, _ := regexp.Compile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()!@:%_\+.~#?&\/\/=]*)`)
	filename := `C:\Users\Dell\go\src\Go_exercise\url_extract\usrls.txt`
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	content := bufio.NewReader(f)
	defer f.Close()

	for {
		line, err := content.ReadBytes('\n')
		if err != nil {
			if err == io.EOF { //必须要io.EOF判断，防止循环EOF报错
				break
			}
			fmt.Println(err)
		}
		url := r.FindString(string(line))
		fmt.Println("URL:", url, "   ", "Domain:", Gethost(url))
	}
}
