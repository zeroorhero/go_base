package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 方法一： 最简单的方法

// 方法二： 类似于while的方法
func main3() {
	i := 2
	for i < 4 {
		fmt.Println("hello")
		i++
	}
}

// 方法三： 类似于死循环的方法，通常搭配break进行使用
// 在for循环开始的时候就进行相关条件的判断
func main4() {
	i := 0
	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i++
	}
}

// 方法四： 使用for range进行判断
// 使用字符串进行举例子
func main5() {
	str := "hello 卢长奇！"
	for i := 0; i < len(str); i++ {
		// 出现乱码的情况的
		// h e l l oåéå
		fmt.Printf("%c \n", str[i])
	}
	fmt.Println("---------------------------")
	// 解决出现乱码的方法：
	// 方法一： 转化为rune[]
	rune_str := []rune(str)
	for i := 0; i < len(rune_str); i++ {
		fmt.Printf("%c \n", rune_str[i])
	}
	// 方法二： 使用for range进行遍历
	for _, val := range str {
		fmt.Printf("%c \n", val)
	}
}

// while 和 do while 的实现

// while的实现
func main6() {
	i := 0
	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i++
	}
}

// do while的实现
func main7() {
	i := 0
	for {
		fmt.Println(i)
		i++
		if i >= 10 {
			break
		}
	}
}

// break的使用

func main() {
	count := 0
	for {
		rand.Seed(time.Now().UnixNano())
		intn := rand.Intn(100)
		fmt.Println(intn)
		if intn == 99 {
			break
		}
		count++
	}
	fmt.Println(count)
}
