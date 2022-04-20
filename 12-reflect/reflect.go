package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

//如果这里写 user *User，则在反射的 inputType.Method()中找不到该方法
func (user User) Call() {
	fmt.Println("user is called...")
	fmt.Printf("%v\n", user)
}

func (user *User) ShowName() {
	fmt.Println("new Name = ", user.Name)
}

func DoFiledAndMethod(input interface{}) {
	//获取 input 的 type
	inputType := reflect.TypeOf(input)
	fmt.Println("type	: ", inputType.Name())

	//获取 input 的 value
	inputValue := reflect.ValueOf(input)
	fmt.Println("value	: ", inputValue)

	//通过 type 获取里面的 字段
	//1. 获取interface的refletc.Type, 通过Type得到字段的数量numField, 进行遍历
	//2. 得到每个 field 数据类型
	//3. 通过filed有一个interface() 方法得到对应的value

	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Printf("%s	: %v = %v\n", field.Name, field.Type, value)
	}

	//通过 type 获取里面的 方法, 调用
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s	: %v\n", m.Name, m.Type)
	}
}

func reflectNum(arg interface{}) {
	fmt.Println("type	: ", reflect.TypeOf(arg))
	fmt.Println("value	: ", reflect.ValueOf(arg))
}

func main() {
	var num float64 = 1.2345
	reflectNum(num)

	fmt.Println("===============================")

	user := User{1, "AceId", 18}
	DoFiledAndMethod(user) //不能获取指针类型的方法

	fmt.Println("===============================")

	inputTypeP := reflect.TypeOf(&user) //可以获取指针类型的方法
	for i := 0; i < inputTypeP.NumMethod(); i++ {
		m := inputTypeP.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}
