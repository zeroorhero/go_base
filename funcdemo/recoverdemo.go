package main

import "fmt"

func test1() {
	// 错误的捕获一定要放在错误可能发生的前面的
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("发生了错误")
		}
	}()
	m := 0
	n := 9
	f := n / m
	fmt.Println(f)
}

func main5() {
	test1()
}
