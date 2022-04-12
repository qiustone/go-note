package main

import "fmt"

func printSlice(mySlice []int) {
	//引用传递
	//_ 表示匿名的变量
	for _, value := range mySlice {
		fmt.Println("value = ", value)
	}
	mySlice[0] = 100
}

func main() {
	myArray := [4]int{1, 2, 3, 4}
	mySlice := []int{1, 2, 3, 4} //动态数据, 切片slice

	fmt.Printf("mySlice type = %T\n", mySlice)
	fmt.Printf("myArray type = %T\n", myArray)

	printSlice(mySlice) //引用传递

	fmt.Println("======================")

	for _, value := range mySlice {
		fmt.Println("value = ", value)
	}

}
