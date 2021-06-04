package main

import (
	"base_go_study/src/lcw/test"
)

func main() {
	println("我是主进程")
	test.PackTest()

	test.StructTest()
	println("我是主进程")
}
