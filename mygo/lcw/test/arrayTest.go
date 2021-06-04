package test

import "fmt"

/**
数组与切片
 */
func Array() {

	//声明一个数组
	var arr [5]int
	arr[0]=1
	//遍历方式一
	for i := range arr {
		fmt.Printf("%v\n",arr[i])
	}
	//遍历方式二
	for i:=0; i<len(arr) ;i++{
		arr[i]=i*2
	}
	fmt.Printf("%v\n",arr)


	//值传递
	arrTwo := arr //等同与  var arrTwo=new([5]int);
	arrTwo[2]=100
	fmt.Printf("arrTwo: %v\n",arrTwo)
	fmt.Printf("arr: %v\n",arr)

	//引用传递 arrThree 修改会把 arr 改掉
	arrThree :=&arr
	arrThree[2]=200
	fmt.Printf("arrThree: %v\n",arrThree)
	fmt.Printf("arr: %v\n",arr)

	//使用数组常量初始化数组  ...表示许多值
	var  arrAge = [...]int {1,2,3,4,5}
	fmt.Printf("arrAge: %v\n",arrAge)

	//有默认大小
	var arrFlow =[5]int {1,2,3}
	fmt.Printf("arrFlow: %v\n",arrFlow)

	//本质切片
	var arrHigh=[]int{100,200,300}
	fmt.Printf("arrHigh: %v\n",arrHigh)

	//指定位置赋值
	var arrWegiht =[5]int {3:2,4:2} //其他数据类型类似
	fmt.Printf("arrWegiht: %v\n\n",arrWegiht)

	//slice
	fmt.Println("slice--------------->>slice\n\n")




}
