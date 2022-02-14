package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello() {
	defer wg.Done()
	fmt.Println("hello lcq!")
}

func main2() {
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go hello()
	}
	fmt.Println("hello main!")
	wg.Wait()
}
