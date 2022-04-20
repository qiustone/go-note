package main

import (
	"fmt"
	"time"
)

//子 goroutine
func newTask() {
	i := 0

	for {
		i++
		fmt.Printf("new Goroutine	: i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

//主 goroutine, 主goroutine退出 子goroutine也会退出
func main() {
	//创建一个go程 去执行 newTask()
	go newTask()

	k := 0
	for {
		k++
		fmt.Printf("main Goroutine	: k = %d\n", k)
		time.Sleep(1 * time.Second)
	}
}
