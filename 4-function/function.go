package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("========== foo1 ==========")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100

	return c
}

//返回多个返回值, 匿名的
func foo2(a string, b int) (int, int) {
	fmt.Println("========== foo2 ==========")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	return 100, 200
}

//返回多个返回值, 有形参名称的
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("========== foo3 ==========")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	//r1 r2 属于foo3的形参, 初始化默认的值是 0
	//r1 r2 作用域空间是 foo3 的整个函数体{}
	fmt.Println("r1 = ", r1, " r2 = ", r2)

	r1 = 300
	r2 = 400
	return
}

func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("========== foo4 ==========")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	r1 = 500
	r2 = 600
	return 700, 800
}

func main() {
	c := foo1("foo1", 111)
	fmt.Println("c = ", c)

	ret1, ret2 := foo2("foo2", 222)
	fmt.Println("ret1 = ", ret1, " ret2 = ", ret2)

	ret1, ret2 = foo3("foo3", 333)
	fmt.Println("ret1 = ", ret1, " ret2 = ", ret2)

	ret1, ret2 = foo4("foo4", 444)
	fmt.Println("ret1 = ", ret1, " ret2 = ", ret2)
}
