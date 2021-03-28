package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/test", ExampleHandler)
	if err := http.ListenAndServe(":3045", nil); err != nil {
		log.Fatal(err)
	}
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {

	// 检查是否为post请求
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "invalid_http_method")
		return
	}

	//application/json格式
	var target map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &target)
	fmt.Println(target["url"])
	defer r.Body.Close()

	fmt.Fprintf(w, "application/json ->"+string(body))
}
