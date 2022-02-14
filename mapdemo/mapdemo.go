package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[int]int{
		3: 89,
		1: 9,
		6: 6,
	}
	var keys []int
	for key, _ := range m {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, i := range keys {
		fmt.Println(m[i])
	}
}
