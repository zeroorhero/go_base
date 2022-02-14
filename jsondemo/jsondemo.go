package main

import (
	"encoding/json"
	"fmt"
)

type animal struct {
	Name string
	Age  int
}

func main() {
	ani := animal{
		Name: "lcq",
		Age:  23,
	}
	marshal, err := json.Marshal(ani)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(string(marshal))
}
