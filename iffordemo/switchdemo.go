package main

import "fmt"

func main1() {
	var input byte
	fmt.Println("请输入一个字母")
	fmt.Scanf("input: %c", &input)
	fmt.Println(input)
	switch input {
	case 'a':
		fmt.Println("今天是星期一")
		fmt.Println("该去上学了")
	case 'b':
		fmt.Println("今天是星期二")
	case 'c':
		fmt.Println("今天是星期三")
	default:
		fmt.Println("今天是周末")
	}
}

func main2() {
	var month byte
	fmt.Println("请输入具体的月份")
	fmt.Scanln(&month)
	switch month {
	case 1, 2, 3:
		fmt.Println("spring")
	case 4, 5, 6:
		fmt.Println("summer")
	default:
		fmt.Println("winter")
	}
}
