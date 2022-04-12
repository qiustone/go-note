package main

import "fmt"

func main() {
	//固定长度的数组
	var myArray1 [10]int

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
}