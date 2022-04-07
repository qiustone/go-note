package main

import "fmt"

//const 来定义枚举类型
const (
	//可以在const() 添加一个关键字 iota, 每行的iota都会累加1，第一行的iota的默认值是0
	BEIJING  = 10 * iota //iota = 0
	SHANGHAI             //iota = 1
	SHENZHEN             //iota = 2
)

const (
	a, b = iota + 1, iota + 2 //iota = 0, a = iota + 1, b = iota + 2, a = 1, b = 2
	c, d                      //iota = 1, c = iota + 1, d = iota + 2, c = 2, d = 3
	e, f                      //iota = 2, e = iota + 1, f = iota + 2, e = 3, f = 4

	g, h = iota * 2, iota * 3 //iota = 3, g = iota * 2, h = iota * 3, g = 6, h = 9
	i, j                      //iota = 4, i = iota * 2, j = iota * 3, i = 8, j = 12
)

func main() {
	//常量(只读属性，不允许修改)
	const length int = 10
	fmt.Println("length = ", length)

	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("SHENZHEN = ", SHENZHEN)

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)
	fmt.Println("e = ", e)
	fmt.Println("f = ", f)
	fmt.Println("g = ", g)
	fmt.Println("h = ", h)
	fmt.Println("i = ", i)
	fmt.Println("j = ", j)

	//iota 只能够配合 const() 一起使用
	//var a int = iota
}
