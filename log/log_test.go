package log_test

import (
	"log"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	log.Println("This is a log.")
	a, b := 10, 20

	log.Print("Use Print to log.")
	log.Println("Ditto for Println.")
	log.Printf("Use Printf and format strings. %d + %d = %d", a, b, a+b)
}

func TestLogformat(t *testing.T) {
	a, b := 10, 20

	// New logger will output to stdout with prefix "Log1: " and flags
	// Note the space in prefix
	myLog := log.New(os.Stdout, "Log1: ", log.Ldate|log.Lshortfile)

	myLog.Print("Use Print to log.")
	myLog.Println("Ditto for Println.")
	myLog.Printf("Use Printf and format strings. %d + %d = %d", a, b, a+b)
}

func TestLogfile(t *testing.T) {
	logFile, err := os.Create("log.txt")
	if err != nil {
		panic("Could not open file.")
	}
	defer logFile.Close()

	a, b := 10, 20

	logmsg := log.New(logFile, "Log:", 0)
	logmsg.Println("Log file create.")
	logmsg.Printf("Use Printf and format strings. %d + %d = %d", a, b, a+b)

	logmsg2 := log.New(logFile, "Log2:", log.Ldate|log.Lshortfile)
	logmsg2.Println("Log file create.")
	logmsg2.Printf("Use Printf and format strings. %d + %d = %d", a, b, a+b)

}
