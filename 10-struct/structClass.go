package main

import "fmt"

//golang 中的类就是使用结构体来绑定方法
//类名、属性名、方法名首字母大写表示对外（其他包）可以访问，否则只能在本包内访问
//如果类名首字母大写, 表示其他包也能够访问
type Hero struct {
	//如果类的属性首字母大写, 表示公有属性，对外能够访问，否则只能在类的内部访问
	Name  string
	Ad    int
	Level int
}

/*
func (this Hero) Show() {
	fmt.Println("Hero	= ", this)
	fmt.Println("Name	= ", this.Name)
	fmt.Println("Ad	= ", this.Ad)
	fmt.Println("Level	= ", this.Level)
}

func (this Hero) GetName() string {
	return this.Name
}

func (this Hero) SetName(newName string) {
	//this 是调用该方法的对象的一个拷贝，调用该方法无法改变该对象该属性的值
	this.Name = newName
}
*/

//首字母大写,表示公有方法
func (this *Hero) Show() {
	fmt.Println("Hero	= ", this)
	fmt.Println("Name	= ", this.Name)
	fmt.Println("Ad	= ", this.Ad)
	fmt.Println("Level	= ", this.Level)
}

func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(newName string) {
	//this 是调用该方法的对象的一个拷贝
	this.Name = newName
}

func main() {
	//创建一个对象
	hero := Hero{Name: "zhang3", Ad: 100, Level: 1}

	hero.Show()

	hero.SetName("Li4")

	hero.Show()
}
