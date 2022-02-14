package main

import "fmt"

type animal struct {
	name string
	age  int
}

type person struct {
	age   int
	name  string
	ptr   *int
	slice []int
	m     map[string]int
}

func main() {
	// 这个是和数组一样的
	// var 定义一个变量就有默认值了
	var a animal
	fmt.Println(a)
	var b person = person{}
	fmt.Println(b)
	fmt.Println(b.m == nil)
	c := &animal{
		name: "lcq",
		age:  0,
	}
	fmt.Println(*c)

}
