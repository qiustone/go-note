package main

import "fmt"

func swap1(a int, b int) {
	var temp int
	temp = a
	a = b
	b = temp
}

func swap2(pa *int, pb *int) {
	var temp int
	temp = *pa
	*pa = *pb
	*pb = temp
}

func main() {
	var a int = 10
	var b int = 20
	swap1(a, b)
	fmt.Println("a = ", a, " b = ", b)
	swap2(&a, &b)
	fmt.Println("a = ", a, " b = ", b)

	var p *int
	p = &a
	fmt.Println(&a)
	fmt.Println(p)

	var pp **int //二级指针
	pp = &p
	fmt.Println(&p)
	fmt.Println(pp)
	fmt.Println(**pp)
}
