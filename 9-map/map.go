package main

import "fmt"

func main() {
	//声明map类型, key和value都是string类型
	var myMap map[string]string

	if myMap == nil {
		fmt.Println("myMap 是一个空 map")
	}

	//在使用map前, 需要先用make给map分配数据空间
	myMap = make(map[string]string, 10)

	myMap["one"] = "java"
	myMap["two"] = "c++"
	myMap["three"] = "python"

	fmt.Println(myMap)
}
