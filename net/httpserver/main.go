package main

import (
	"log"
	"net/http"
)
//Fileserver
func main(){
	err := http.ListenAndServe(":80",http.FileServer(http.Dir(".")))
	if err != nil{
		log.Fatal(err)
	}
}
