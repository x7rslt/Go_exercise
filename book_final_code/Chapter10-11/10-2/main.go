package main

import "fmt"
import "sync"

var wg sync.WaitGroup

func main() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}

	wg.Wait()
}
