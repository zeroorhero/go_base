package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main1() {
	// 定义byte实际上是一个数字的
	var input byte
	fmt.Println("请输入一个字母")
	scan, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(scan)
	fmt.Printf("the anser is %c", input)
}

func main2() {
	var input byte
	fmt.Println("请输入一个字母")
	_, err := fmt.Scanf("%c", &input)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Printf("the anser is %c", input)
}

func main3() {
	var input string
	fmt.Println("请输入字母")
	fmt.Scan(&input)
	fmt.Println(input)
}

func main4() {
	var input string
	fmt.Println("请输入信息")
	fmt.Scanf("input is %s", &input)
	fmt.Println(input)
}

func main5() {
	var (
		name    string
		age     int
		married bool
	)
	fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}

func main6() {
	var (
		name    string
		age     int
		married bool
	)
	fmt.Scanln(&name, &age, &married)
	fmt.Println(name, age, married)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入内容")
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(strings.TrimSpace(str))
}
