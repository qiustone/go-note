package main

import "fmt"

type myint int //myint 是 int 的一个别名

//定义一个结构体
type Book struct {
	title string
	auth  string
}

func fakeChangeBook(book Book) {
	//传递的是book的副本, 本修改不生效
	book.auth = "666"
}

func changeBook(book *Book) {
	//引用传递(指针传递)
	book.auth = "777"
}

func main() {
	/*
		var a myint = 10
		fmt.Println("a = ", a)
		fmt.Printf("type of a = %T", a) //type of a = main.myint
	*/

	//结构体赋值
	var book1 Book
	book1.title = "Golang"
	book1.auth = "zhang3"
	fmt.Printf("%v\n", book1)

	fakeChangeBook(book1)
	fmt.Printf("%v\n", book1)

	changeBook(&book1)
	fmt.Printf("%v\n", book1)
}
