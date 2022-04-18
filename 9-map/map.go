package main

import "fmt"

func printMap(cityMap map[string]string) {
	//cityMap 是一个引用传递
	fmt.Println("=====================================")
	for key, value := range cityMap {
		fmt.Println("Key = ", key, "	val = ", value)
	}
}

func changeValue(cityMap map[string]string) {
	cityMap["England"] = "London"
}

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

	fmt.Println("=====================================")

	//map的使用
	cityMap := make(map[string]string)

	//添加
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	//遍历
	for key, value := range cityMap {
		fmt.Println("Key = ", key, "	val = ", value)
	}

	//删除
	delete(cityMap, "Japan")

	//修改
	cityMap["USA"] = "DC"

	//默认是引用传递
	changeValue(cityMap)

	//传参
	printMap(cityMap)
}
