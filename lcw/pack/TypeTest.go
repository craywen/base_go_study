package pack

import (
	"fmt"
)

type Books struct {
	title   string
	author  string
	subject string
	bookId  int
}

func TypeTest() {
	// 创建一个新的结构体
	b := new(Books)
	b.author = "java"
	b.bookId = 123
	b.subject = "语文"
	b.title = "阿斯达"
	fmt.Println(b)
	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", bookId: 6495407})

	// 忽略的字段为 0 或 空
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})
	c := new(Books)
	_, _ = getBooks(c)
	fmt.Println(c)

	var slice1 []int = make([]int, 10)
	slice1 = append(slice1, 1)
	fmt.Println(slice1)
}

func getBooks(books *Books) (a Books, c int) {
	books.bookId = 12312312312312313
	fmt.Println(books)
	return *books, 1
}
