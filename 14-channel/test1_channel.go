package main

import "fmt"

func main() {
	/*
		channel的定义
		make(chan Type)   //等价于 make(chan Type, 0)
		make(chan Type, capacity)

		channel <- value   //发送value到channel
		<-channel          //接收并将其丢弃
		x := <-channel     //从channel中接收数据, 并赋值给 x
		x, ok := <-channel //功能同上, 同时检查通道是否已关闭或者是否为空
	*/

	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine结束")

		fmt.Println("goroutine 正在运行")

		c <- 666 //将666发送给c
	}()

	num := <-c //从c中接收数据，并赋值给num

	fmt.Println("num = ", num)
	fmt.Println("main goroutine 结束...")
}
