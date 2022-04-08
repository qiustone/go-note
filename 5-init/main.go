package main

//导包需从$GOPATH开始
import (
	"go-note/5-init/lib1"
	"go-note/5-init/lib2"
)

//执行顺序 import lib1 -> lib1.const -> lib1.var -> lib1.init() -> main()
func main() {
	lib1.Lib1Test()
	lib2.Lib2Test()
}
