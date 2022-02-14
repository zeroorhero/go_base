package main

import (
	"fmt"
)

func main10() {
	wg.Add(2)
	write := make(chan int, 5)
	read := make(chan int, 5)
	go writechan(write)
	go readchan(write, read)
	for val := range read {
		fmt.Printf("main 为%d \n", val)
	}
	wg.Wait()
}

func writechan(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
	wg.Done()
}

func readchan(write, read chan int) {
	for val := range write {
		fmt.Printf("读出的数据为%d \n", val)
		read <- val
	}
	close(read)
	wg.Done()
}
