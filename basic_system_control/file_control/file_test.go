package file_test

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestCreate(t *testing.T) {
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}

func TestDelete(t *testing.T) {
	err := os.Remove("test.txt")
	if err != nil {
		log.Fatal(err)
	}

}

type People struct {
	name string
	age  int
}

func TestExist(t *testing.T) {
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("File does exist.File information:")
	log.Println(fileInfo)
}

func TestFileInfo(t *testing.T) {
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory:", fileInfo.IsDir())
	fmt.Printf("Sytem interface type:%T\n", fileInfo.Sys())
	fmt.Printf("System info:%+v\n\n", fileInfo.Sys())

	xiao := People{
		"xiaoyang",
		14,
	}
	fmt.Printf("People :%v\n", xiao)
	fmt.Printf("People :%+v", xiao)

}

func TestMove(t *testing.T) {
	oldpath := "test.txt"
	newpath := "test2.txt"

	err := os.Rename(oldpath, newpath)

	if err != nil {
		log.Fatal(err)
	}
}

func TestT(t *testing.T) {
	// Truncate a file to 100 bytes. If file
	// is less than 100 bytes the original contents will remain
	// at the beginning, and the rest of the space is
	// filled will null bytes. If it is over 100 bytes,
	// Everything past 100 bytes will be lost. Either way
	// we will end up with exactly 100 bytes.
	// Pass in 0 to truncate to a completely empty file

	err := os.Truncate("test2.txt", 10)
	if err != nil {
		log.Fatal(err)
	}
}
