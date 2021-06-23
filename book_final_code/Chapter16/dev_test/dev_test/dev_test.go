package dev_test

import (
	"fmt"
	"testing"
)

func TestDev(t *testing.T) {
	fmt.Println("hello")
	Hello()
}

func Hello() {
	fmt.Println("This is a test.")
}
