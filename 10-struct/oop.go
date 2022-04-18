package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

type SuperMan struct {
	Human //SuperMan 类继承了Human 类的方法
	level int
}

//重定义父类的方法Eat()
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

//子类的新方法
func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func main() {
	h := Human{"zhang3", "female"}
	h.Eat()
	h.Walk()

	fmt.Println("=======================")

	//定义一个子类的对象
	//第一种方式
	//s := SuperMan{Human{"li4", "female"}, 88}
	//第二种方式
	var s SuperMan
	s.name = "li4"
	s.sex = "female"
	s.level = 88

	s.Walk() //父类的方法
	s.Eat()  //子类的方法
	s.Fly()  //子类的方法
}
