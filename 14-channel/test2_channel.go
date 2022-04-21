package main

import (
	"fmt"
	"time"
)

//带缓冲区的channel
//当channel已满, 再向里面写数据, 就会阻塞
//当channel为空, 从里面取数据也会阻塞

func main() {
	c := make(chan int, 3) //带缓冲的channel

	fmt.Println("len(c) = ", len(c), ", cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("子go程结束")

		//当发送的元素大于缓冲区大小, 而接收方还未接收时，会发生阻塞
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("子go程正在运行,发送的元素 = ", i, " len(c) = ", len(c), ", cap(c) = ", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c //从 channel 中接收数据, 并赋值给 num
		fmt.Println("num = ", num)
	}

	time.Sleep(2 * time.Second) //防止子go程提前结束

	fmt.Println("main 结束")
}
