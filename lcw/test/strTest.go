package test

import (
	"fmt"
	"strconv"
	"strings"
)

func StrTest() {

	var name string = "wodehasw,asdasd,qweqwe "
	var fix bool = strings.HasPrefix(name, "f")
	var fixend bool = strings.HasSuffix(name, "e")
	fmt.Printf("%t\n", fix)
	fmt.Printf("%t\n", fixend)

	//包含关系
	has := strings.Contains(name, "od")
	fmt.Printf("%t\n", has)
	//index
	index := strings.Index(name, "d")
	fmt.Printf("%v\n", index)
	//替换
	replace := strings.Replace(name, "wo", "no", 1)
	fmt.Printf("%s\n", replace)
	//次数
	count := strings.Count(name, "w")
	fmt.Printf("count %v\n", count)
	//去空格
	newStr := strings.TrimSpace(name)
	fmt.Printf("%v\n", newStr)
	//字符串遍历
	sl2 := strings.Split(name, ",")
	for _, value := range sl2 {
		fmt.Printf("%v\n", value)
	}
	//字符串间隔拼接
	newStr2 := strings.Join(sl2, "lcw ")
	fmt.Printf("newStr2  %v\n", newStr2)
	//字符串byte
	newpool := strings.NewReader(name)
	i := newpool.Len()
	fmt.Printf("newpool  %v  and %d\n", newpool, i)
	//字符串转换
	var str string = strconv.Itoa(i)
	fmt.Printf("str:  %v\n", str)

	val, err := strconv.Atoi(str)
	fmt.Printf("val%v\t %v\n", val, err)

}
