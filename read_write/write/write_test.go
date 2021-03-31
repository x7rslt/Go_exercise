package write_test

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestQuickWrite(t *testing.T) {
	err := ioutil.WriteFile("test.txt", []byte("Helllo world.\n"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func TestWriteBytes(t *testing.T) {
	file, err := os.OpenFile("test2.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	byteSlice := []byte("Bytes!跳舞\n")
	bytesWrite, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWrite)
}
