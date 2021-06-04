package test

import (
	"fmt"
	"strconv"
)

func FuncTest() {
	//一般使用多参数返回,正常情况下,默认返回 nil
	val,err :=strconv.Atoi("12312313");
	if err ==nil{
		fmt.Printf("%v ,%v\n",val,err)
	}else{
		fmt.Printf("error %v ,%v\n",val,err)
	}
	//defer 关键词,相当于java c#的 finally
	/*
	关键字 defer 允许我们进行一些函数执行完成后的收尾工作，例如：
	关闭文件流：
	// open a file defer file.Close() （详见第 12.2 节）
	解锁一个加锁的资源
	mu.Lock() defer mu.Unlock() （详见第 9.3 节）
	打印最终报告
	printHeader() defer printFooter()
	关闭数据库链接
	// open a database connection defer disconnectFromDB()
	合理使用 defer 语句能够使得代码更加简洁。
	*/
	defer aa(10)

	s	:=bb("20");
	fmt.Printf("%s\n",s)
}

func aa(a int ) int {
	fmt.Printf("defer  %v\n",a)
	return 10;
}

func bb(a string) string {
	return  "20"
}

func aaa()  {
	
}
