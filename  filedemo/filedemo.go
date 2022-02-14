package main

import (
	"fmt"
	"os"
)

func main(){
	file, err := os.Open("/Users/lcq/Desktop/a.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

}
