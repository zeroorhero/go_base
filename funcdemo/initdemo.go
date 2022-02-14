package main

import "fmt"

var age = test()

func test() int {
	fmt.Println("test")
	return 90
}

func init() {
	fmt.Println("init start")
}

func main3() {
	fmt.Println("main start")
}
