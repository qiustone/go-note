package main

import "fmt"

func main() {
	//============= 第一种声明方式 ================
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

	//============= 第二种声明方式 ================
	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "c++"
	myMap2[3] = "python"

	fmt.Println(myMap2)

	//============= 第三种声明方式 ================
	myMap3 := map[string]string{
		"one":   "php",
		"two":   "c++",
		"three": "python",
	}
	fmt.Println(myMap3)
}
