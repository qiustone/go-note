package main

import "fmt"

/*
	单流程下一个 go 只能监听一个 channel 状态, select 可以完成监听多个 channel 状态
	配合for循环使用可持续监听
*/

/*
	select {
	case <-chan1:
		//如果chan1成功读取到数据，则进行该 case 处理语句
	case chan2 <- l:
		//如果成功向chan2写入数据，则进行该 case 处理语句
	default:
		//如果上面都没有成功, 则进入 default 处理流程
	}
*/

func fibonacii(c, quit chan int) {
	x, y := 1, 1

	for {
		select {
		case c <- x: //如果 c 可写
			x = y
			y = x + y
		case <-quit: //如果 quit 可读
			fmt.Println("quit中有数据可读, 退出主程序")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	//sub go
	go func() {
		for i := 1; i <= 6; i++ {
			fmt.Println(i, "读取c中的数据: ", <-c) //如果管道 c 可读, 打印 c 中的数据
		}

		quit <- 0 //将数据写入quit
	}()

	//mai go
	fibonacii(c, quit)
}
