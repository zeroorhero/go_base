package main

import "fmt"

func main() {
	// 方法一：
	var a [3]int = [3]int{2, 3, 4}
	fmt.Println(a)

	// 方法二：
	var b = [3]int{4, 5, 6}
	fmt.Println(b)

	// 方法三：
	c := [...]int{8, 9}
	fmt.Println(c)

	// 方法四：
	d := [...]int{1: 8, 4: 3}
	fmt.Println(d)

	// 数组不用进行初始化 直接就是该类型的默认值了
	var e [4]int
	fmt.Println(e)
}
