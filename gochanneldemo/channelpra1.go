package main

import (
	"fmt"
)

// 从现有的通道中取数据】
// 方法一： for 死循环
// 方法二： for range
func main12() {
	data := make(chan int, 1000)
	cal := make(chan int, 2000)
	wg.Add(1)
	go Putdata(data)
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go Getdata(data, cal)
	}
	wg.Wait()
	close(cal)
	for val := range cal {
		fmt.Println(val)
	}
}

func Putdata(data chan int) {
	defer wg.Done()
	for i := 1; i <= 8000; i++ {
		data <- i
	}
	close(data)
}

func Getdata(data, cal chan int) {
	defer wg.Done()
	for num := range data {
		flag := true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			cal <- num
		}
	}
}
