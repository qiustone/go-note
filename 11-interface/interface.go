package main

import "fmt"

//定义一个接口，本质是一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string //获取动物的颜色
	GetType() string  //获取动物的种类
}

//具体的类, 只要实现了AnimalIF接口的所有方法, 就是继承了该接口, 仅实现部分方法不能实现该接口
type Cat struct {
	color string //猫的颜色
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is Sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is Sleep")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AnimalIF) {
	animal.Sleep()
	fmt.Println("color	= ", animal.GetColor())
	fmt.Println("kind	= ", animal.GetType())
}

//interface{} 是万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)

	//interface{} 如何区分, 此时应用的底层数据类型到底是什么？

	//Golang 给 interface 提供 "类型断言"的机制
	val, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type, value = ", val)
		fmt.Printf("value type is %T\n", val)
	}
	fmt.Println("=================================")
}

type Book struct {
	auth string
}

/*
多态:
1. 基本要素  :
	有一个父类(接口)
	有子类(实现了父类的全部接口方法)
	父类类的变量(指针) 指向(引用) 子类的具体数据变量

2. interface
	通用万能类型 interface{} 空接口
	int, string, float32, float64, struct 都实现类 interface{}
	可以用 interface{} 类型来引用
*/

func main() {
	/*
		var animal AnimalIF //接口的数据类型，父类指针
		animal = &Cat{"Green"}  //接口本质是一种指针, 所以需要将对象的地址赋值给接口
		animal.Sleep()

		animal = &Dog{"Yellow"}
		animal.Sleep()
	*/

	cat := Cat{"Green"}
	dog := Dog{"Yellow"}

	showAnimal(&cat) //接口本质是一种指针, 所以需要将对象的地址赋值给接口
	showAnimal(&dog)

	fmt.Println("=================================")

	book := Book{"Golang"}
	myFunc(book)
	myFunc(100)
	myFunc("ABC")
	myFunc(3.14)
}
