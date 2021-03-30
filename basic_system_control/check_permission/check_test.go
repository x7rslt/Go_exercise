package check_test

import (
	"log"
	"os"
	"testing"
)

func TestCheck(t *testing.T) {
	file, err := os.OpenFile("test.txt", os.O_WRONLY, 0000)
	if err != nil {
		if !os.IsPermission(err) {
			log.Printf("Write error:%s", err)
		}
		log.Println("Error:Write Permission denied.")

	}
	file.Close()

	file, err = os.OpenFile("test.txt", os.O_RDONLY, 0000)
	if err != nil {
		if !os.IsPermission(err) {
			log.Printf("Read error:%s", err)
		}
		log.Println("Error:Read Permission denied.")
	}
	file.Close()
}
