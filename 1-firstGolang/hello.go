//程序的包名，当前程序所在的包
//如果当前文件中有main函数，则这个包一定是main包
//跟当前的文件名hello.go没有任何的关系
package main

import "fmt"

//main 函数
func main() { //函数的左花括号必须与函数名在同一行，否则编译错误
	//golang 中的表达式, 加不加“;”都可以, 建议不加
	fmt.Println("Hello Go !")
}
