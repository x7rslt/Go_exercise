package read_test

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestReadAll(t *testing.T) {
	file, err := os.Open("dict.txt")
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data as hex:%x\n", data)
	fmt.Printf("Data as string:%s\n", data)
	fmt.Printf("Number of bytes read:%d", len(data))

}

//逐行读取.csv文件
func TestBufioread(t *testing.T) {
	reader := bufio.NewReader(os.Stdin) //不适合test，放在main函数；go run后，输入任意一行字符即返回打印。
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(input)
}

func TestCsv(t *testing.T) {
	//准备读取文件
	fileName := "./top.csv"
	fs, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
	}
	defer fs.Close()

	r := csv.NewReader(fs)
	//针对大文件，一行一行的读取文件
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		fmt.Println(row[1])
	}
}

func TestReadScan(t *testing.T) {
	wordfile, err := os.Open(`dict.txt`)
	if err != nil {
		log.Fatal(err)
	}
	defer wordfile.Close()

	done := make(chan bool)
	scanner := bufio.NewScanner(wordfile)
	go func() {
		for scanner.Scan() {
			pass := scanner.Text()
			fmt.Println(pass)
		}
		done <- true
	}()
	<-done
}
