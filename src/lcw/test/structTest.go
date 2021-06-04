package test

import "fmt"

type Man struct {
	name string
	age  int
}

func StructTest() {
	// 使用一个结构体
	man := new(Man)
	man.name = "lcw"
	man.age = 12
	fmt.Printf("man %v", man)
}
