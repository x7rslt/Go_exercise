package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	tpl, err := template.ParseGlob("/Users/xiaoshuai/go/src/Go_exercise/html_template/html/view/login.html")
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range tpl.Templates() {
		tplname := v.Name()
		fmt.Println(v.Name())
		http.HandleFunc(tplname, func(writer http.ResponseWriter, request *http.Request) {
			_ = tpl.ExecuteTemplate(writer, tplname, nil)
		})
	}
}
