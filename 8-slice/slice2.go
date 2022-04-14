package main

import "fmt"

func main() {
	//1. slice 的四种声明方式
	//第一种: 声明slice1 是一个切片, 并且初始化, 值是1,2,3 长度len是3
	slice1 := []int{1, 2, 3}
	fmt.Printf("len = %d, slice1 = %v\n", len(slice1), slice1)

	//第二种: 声明slice2 是一个切片, 但是并没有给slice2 分配空间, 后面再分配空间
	var slice2 []int
	fmt.Printf("len = %d, slice2 = %v\n", len(slice2), slice2)

	//slice2[0] = 1 //此时给slice2赋值将会报错, 因为没有给slice2 分配空间

	slice2 = make([]int, 3) //给slice2 分配空间
	fmt.Printf("len = %d, slice2 = %v\n", len(slice2), slice2)

	//第三种:声明slice3 是一个切片, 同时给slice3 分配空间, 初始化值是0
	var slice3 []int = make([]int, 3)
	fmt.Printf("len = %d, slice3 = %v\n", len(slice3), slice3)

	//第四种: 属于第三种的简写
	slice4 := make([]int, 3)
	fmt.Printf("len = %d, slice4 = %v\n", len(slice4), slice4)

	//判断slice是否为空
	var slice5 []int
	if slice5 == nil {
		fmt.Println("slice5 是一个空切片")
	} else {
		fmt.Println("slice5 不是一个空切片")
	}

	fmt.Println("=========================")

	//2. 切片追加
	var numbers = make([]int, 3, 5)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// numbers[4] = 4, 由于长度为3, 这里会报错 runtime error: index out of range [4] with length 3
	// name如何向容量为5的slice(切片)追加元素呢？可以使用append
	numbers = append(numbers, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	//向一个容量cap已满的slice 追加元素, golang 会自动将numbers 的cap 增加一倍,也就是cap=cap*2
	numbers = append(numbers, 2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	//不指定slice容量, 容量默认 = slice的长度
	var numbers2 = make([]int, 3) //len = 3, cap = 3
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)

	numbers2 = append(numbers2, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)

	fmt.Println("=========================")

	//3. 切片的截取和拷贝
	s := []int{1, 2, 3} //len = 3, cap = 3
	s1 := s[0:2]        //相当于数学中的[0:2) = [1, 2]
	fmt.Println(s1)
	s2 := s[:]
	fmt.Println(s2)
	s3 := s[1:]
	fmt.Println(s3)
	//s4 := s[0:-1] 并不支持负数
	//fmt.Println(s4)

	//s1 与 s 的指针指向的是同一位置, 改变s1中的元素也会影响s中的元素
	s1[0] = 100
	fmt.Println("s = ", s)
	fmt.Println("s1 = ", s1)

	//copy 函数可以拷贝slice中的值
	ss := make([]int, 3) //[0,0,0]

	//将s中的值依次拷贝的s2中
	copy(ss, s)

	fmt.Println(ss)
}
