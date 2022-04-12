package main

import "fmt"

func printArray(myArray [4]int) {
	//值拷贝
	for index, value := range myArray {
		fmt.Println("index = ", index, " value = ", value)
	}

	myArray[0] = 111
}

func main() {
	//固定长度的数组
	var myArray1 [10]int //默认值为 0

	myArray2 := [10]int{1, 2, 3, 4}
	myArray3 := [4]int{11, 22, 33, 44}

	for i := 0; i < 10; i++ {
		fmt.Println(myArray1[i])
	}

	for index, value := range myArray2 {
		fmt.Println("index = ", index, ", value = ", value)
	}

	fmt.Printf("myArray1 types = %T\n", myArray1)
	fmt.Printf("myArray2 types = %T\n", myArray2)
	fmt.Printf("myArray2 types = %T\n", myArray3)

	//printArray(myArray1) 因为数据长度也是类型的一部分，所以类型不匹配
	printArray(myArray3)

	//数组传参是值拷贝, 函数内部对数组的修改, 并不会影响外部的数组
	for index, value := range myArray3 {
		fmt.Println("index = ", index, ", value = ", value)
	}

}
