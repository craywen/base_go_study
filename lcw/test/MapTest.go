package test

import (
	"fmt"
	"sort"
	"strconv"
)

func MapTest() {
	// var map1 map[keytype]valuetype
	//keytype key 可以是任意可以用 == 或者 != 操作符比较的类型，比如 string、int、float。所以数组、切片和结构体不能作为 key，但是指针和接口类型可以
	//valuetype value 可以是任意类型的；
	var mapList map[string]int
	mapList = map[string]int{"s": 1, "a": 2}
	fmt.Printf("%v\n", mapList)

	//var makeMap map[string] int
	//使用make 创建
	makeMap := make(map[string]float64)
	makeMap["key1"] = 4.5
	makeMap["key2"] = 3.14159
	makeMap["two"] = 3
	fmt.Printf("makeMap:---> %v\n", makeMap)

	//tips :不要使用 new，永远用 make 来构造 map
	//注意 如果你错误的使用 new() 分配了一个引用对象，你会获得一个空引用的指针，相当于声明了一个未初始化的变量并且取了它的地址：
	//为了说明值可以是任意类型的，这里给出了一个使用 func() int 作为值的 map
	//注意是,分隔而不是回车或者;
	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		3: func() int { return 40 },
	}
	fmt.Printf("mf:---> %v\n", mf)

	//创建map时指定容量 注意 map 不是按照 key 的顺序排列的，也不是按照 value 的序排列的。
	mapCap := make(map[string]int, 100)
	for i := 0; i < 100; i++ {
		mapCap[strconv.Itoa(i)] = i
	}
	fmt.Printf("mapCap:---> %v\n", len(mapCap))
	//遍历map
	for s := range mapCap {
		val := mapCap[s]
		fmt.Println(val)
	}
	//带key value
	for key, value := range mapCap {

		fmt.Printf("%v ,%v\n", key, value)
	}

	//取具体的值
	fmt.Printf("mapCap[\"12\"]:---> %v\n", mapCap["12"])

	//判断map的key是不是存在
	if _, ok := mapCap["100"]; ok {
		fmt.Printf("ok:---> %v\n", ok)
	} else {
		fmt.Printf("ok:---> %v\n", ok)
	}

	//删除key
	delete(mapCap, "99")
	value, res := mapCap["99"]
	if res {
		fmt.Printf("%v\n", value)
	} else {
		fmt.Printf("%v\n", len(mapCap))
	}
	//map Version A 切片 -->简称map 数组
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}
	fmt.Printf("Version A: Value of items: %v\n", items)

	// Version B: NOT GOOD!
	items2 := make([]map[int]int, 5)
	for _, item := range items2 {
		item = make(map[int]int, 1) // item 只是 slice 元素的副本。
		item[1] = 2                 // 这个“项目”将在下一次迭代中丢失。
	}
	fmt.Printf("Version B: Value of items: %v\n", items2)

	//map 的排序
	keys := make([]string, len(mapCap))
	index := 0
	for key, _ := range mapCap {
		keys[index] = key
		index++
	}
	sort.Strings(keys)
	fmt.Printf("keys sort: %v\n", keys)
	for _, val := range keys {
		fmt.Printf("mapCap sort: %v\n", mapCap[val])
	}
	//map键值对调,新建一个map 遍历对调就行
	invMap := make(map[int]string, len(mapCap))
	for key, value := range mapCap {
		invMap[value] = key
	}
	for k, v := range invMap {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}

}
