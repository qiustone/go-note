package lib2

import "fmt"

func init() {
	fmt.Println("lib2.init() run ...")
}

//当前lib2包提供的API，大写函数名表示可供外部调用，小写函数名外部无法调用
func Lib2Test() {
	fmt.Println("lib2Test()...")
}
