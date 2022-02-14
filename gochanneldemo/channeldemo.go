package main

import (
	"fmt"
	"time"
)

var m = make(map[int]int, 10)

func solution(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res = res * i
	}
	m[n] = res
}

func main4() {
	for i := 1; i <= 200; i++ {
		go solution(i)
	}
	time.Sleep(time.Second * 10)
	for ind, val := range m {
		fmt.Printf("[%d] = %d \n", ind, val)
	}
}

func main5() {
	var ch = make(chan int, 2)
	ch <- 2
	ch <- 3
	// ch <- 8
	// 开始取出数据
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// 测试关闭channel
func main6() {
	var ch = make(chan int, 3)
	ch <- 2
	ch <- 3
	// ch <- 8
	close(ch)
	// panic: send on closed channel
	ch <- 7
}

func test(ch chan int) {
	ch <- 4
	close(ch)
	defer wg.Done()
}
func main9() {
	var ch = make(chan int, 3)
	ch <- 2
	ch <- 3
	// close(ch)
	wg.Add(1)
	go test(ch)
	// for range遍历 也就是将channel中的数据取出来的
	for val := range ch {
		fmt.Println(val)
	}
	// fmt.Println(<-ch)
	wg.Wait()
}
func receve(ch chan int) {
	<-ch
	fmt.Println("接受成功")
}

func main7() {
	var ch chan int = make(chan int)
	go receve(ch)
	ch <- 19
	fmt.Println("发送成功")

}
