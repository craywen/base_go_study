package test

import "fmt"

func PointerTest() {
	//声明一个变量,打印它的值和它的内存地址
	var integer int = 9
	fmt.Printf("%v %v\n", integer, &integer)
	//将一个变量的指针指向这个 integer
	var fromToP *int
	fromToP = &integer
	var c = *fromToP
	c = 8
	fmt.Printf("%v %v\n", fromToP, *fromToP)
	fmt.Printf("%v %v\n", c, &c)
	fmt.Printf("%v %v\n", fromToP, *fromToP)

	s := "hollo word"
	var p *string = &s
	*p = "ciao"
	fmt.Printf("here is s %p\n", p)
	fmt.Printf("here is p %s\n", *p)
	fmt.Printf("here is s %s\n", s)

	s1 := "good bye"
	var p1 *string = &s1
	*p1 = "ciao"
	fmt.Printf("Here is the pointer p1: %p\n", p1)  // prints address
	fmt.Printf("Here is the string *p1: %s\n", *p1) // prints string
	fmt.Printf("Here is the string s1: %s\n", s1)   // prints same string

}
