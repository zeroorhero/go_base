package main

import "fmt"

func main() {
	var x interface{}
	var b float32 = 1.1
	x = b
	y := x.(float32)
	fmt.Println(y)
	fmt.Printf("%t", y)
}
