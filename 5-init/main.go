package main

//导包需从$GOPATH开始
import (
	"fmt"
	_ "go-note/5-init/lib1"   //匿名导包，无法使用里面的方法，但是会执行init()函数
	lib "go-note/5-init/lib2" //别名导包
	. "go-note/5-init/lib3"   //将所有函数导入当前包，可以直接调用
)

//执行顺序 import lib1 -> lib1.const -> lib1.var -> lib1.init() -> main()
func main() {
	//lib1.Lib1Test()
	fmt.Println("================")
	lib.Lib2Test()
	Lib3Test()
}
