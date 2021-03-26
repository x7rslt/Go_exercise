package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	err = os.Chmod("test.txt", 0777)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Chown("test.txt", os.Getuid(), os.Getgid())
	if err != nil {
		log.Fatal(err)
	}
}
