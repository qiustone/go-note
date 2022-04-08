package lib1

import "fmt"

func init() {
	fmt.Println("lib1.init() run ...")
}

//当前lib1包提供的API，大写函数名表示可供外部调用，小写函数名外部无法调用
func Lib1Test() {
	fmt.Println("lib1Test()...")
}
