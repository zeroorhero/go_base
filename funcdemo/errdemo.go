package main

import "errors"

func confi(name string) error {
	if "config" == name {
		return nil
	} else {
		return errors.New("发生了错误")
	}
}

// 出现err 不一定程序结素
// 但是要是出现panic的话  程序一定会结束的
func main() {
	err := confi("hhhh")
	if err != nil {
		panic(err)
	}
}
