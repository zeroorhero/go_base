package main

import "fmt"

func getSum(m, n int) int {
	return m + n
}

func sum(a int, args ...int) int {
	sum := a
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func main2() {
	i := sum(2, 3, 4, 6, 7)
	fmt.Println(i)
}

type myfunc func(int, int) int

func main1() {
	// 将一个具体的函数赋值给一个变量了
	// 在这里面省略了a的具体的类型
	// a的类型为func(int, int)int
	a := getSum
	fmt.Println(a(2, 7))

	var b myfunc
	b = getSum
	fmt.Println(b(1, 3))
}
