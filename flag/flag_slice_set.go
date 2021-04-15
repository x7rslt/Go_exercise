package main

import (
	"flag"
	"fmt"
	"log"
)

type Foo []string

func (f *Foo) String() string {
	return ""
}

func (f *Foo) Set(s string) error {
	*f = append(*f, s)
	return nil
}

var foo Foo

func main() {
	flag.Var(&foo, "foo", "foo")
	flag.Parse()
	fmt.Println(foo)
	log.Printf("%#v", foo)
}
