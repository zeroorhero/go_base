package main

import (
	"fmt"
	"time"
)

func main1() {
	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()
	fmt.Println(year, month, day)
}

func main() {
	unix := time.Now().Unix()
	fmt.Println(unix)
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05"))
}
