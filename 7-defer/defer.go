package main

import "fmt"

func func1() {
	fmt.Println("A")
}

func func2() {
	fmt.Println("B")
}

func func3() {
	fmt.Println("C")
}

func defer_call() {
	defer func1()
	defer func2()
	defer func3()
}

func deferFunc() int {
	fmt.Println("defer func called ...")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called ...")
	return 0
}

//return 先调用, defer 后调用
func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	//defer 关键字
	defer fmt.Println("main end 1") //进栈，先进后出
	defer fmt.Println("main end 2") //进栈，后进先出

	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 2")

	fmt.Println("=============================")
	defer_call()
	fmt.Println("=============================")
	returnAndDefer()
	fmt.Println("=============================")
}
