package main

import "fmt"

/*
	变量的四种声明方式
*/

//方法一、方法二、方法三是可以声明全局变量的，但方法四不可以
var gA int = 22
var gB = 220

func main() {
	//方法一：声明一个变量, 默认的值是0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("Type of a = %T\n", a)

	//方法二：声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("Type of b = %T\n", b)

	//方法三：在初始化的时候， 可以省去数据类型，通过值自动匹配当前变量的数据类型
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("Type of c = %T\n", c)

	//方法四：(常用的方法) 省去var关键字，使用 := 自动匹配, := 既初始化又赋值
	d := 100
	fmt.Println("d = ", d)
	fmt.Printf("type of d = %T\n", d)

	e := "abc"
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T\n", e)

	//====== 全局变量
	fmt.Println("gA =", gA, " gB =", gB)

	//声明多个变量
	//1. 相同类型
	var xx, yy int = 100, 200
	fmt.Println("xx =", xx, " yy =", yy)

	//2. 不同类型
	var mm, nn = 100, "Letshi"
	fmt.Println("mm =", mm, " nn =", nn)

	//3. 多行多变量声明
	var (
		kk int  = 100
		jj bool = true
	)
	fmt.Println("kk =", kk, " jj =", jj)
}
