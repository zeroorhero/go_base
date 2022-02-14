package main

import (
	"fmt"
	"runtime"
)

func main1() {
	a := runtime.NumCPU()
	fmt.Println(a)
	runtime.GOMAXPROCS(a)
}
