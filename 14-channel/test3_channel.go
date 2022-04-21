package main

import "fmt"

/*

channel 不像文件一样需要经常关闭, 只有当你确实没有数据发送了，
或者你想显式的结束range循环之类的，才去关闭channel;

关闭channel后，无法向channel再发送数据(引发panic错误后导致接收立即返回零值);

关闭channel后，可以继续从channel接收数据;

对于 nil channel, 无论收发都会被阻塞

*/

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
			//close(c)  //关闭后再发数据会发生错误, panic: send on closed channel
		}

		//close 可以关闭一个channel, 如果不close会形成死锁
		close(c)
	}()

	for {
		//ok 如果为 true 表示channel没有关闭, 如果为false表示channel已经关闭
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("Main Finished...")
}
