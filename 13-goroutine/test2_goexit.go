package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	//用go创建一个匿名go程
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			//退出当前goroutine, 与 return 的区别是, return只能退出当前匿名函数, 无法退出外层匿名函数
			runtime.Goexit()
			//return
			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	//那么如何获取下面匿名函数的返回值呢？需要用到channel
	go func(a int, b int) bool {
		fmt.Println("a = ", a, ", b = ", b)
		return true
	}(10, 20)

	//死循环
	for {
		time.Sleep(1 * time.Second)
	}
}
