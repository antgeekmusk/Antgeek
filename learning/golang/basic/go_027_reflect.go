package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

// user的方法
func (this User) UserCall() {
	fmt.Println("UserCall run\n")
}
func (this User) UserCallWithArgs(a int, b string) {
	fmt.Println("UserCallWithArgs run, args = ", a, b)
}

// 通过反射获取结构体的类型,值,方法 以及其属性的类型值
func doReflect(t interface{}) {
	// 获取t的反射值
	tValue := reflect.ValueOf(t)
	tType := reflect.TypeOf(t)

	// 获取属性
	for i := 0; i < tType.NumField(); i++ {
		field := tType.Field(i)              // 获取属性的类型
		value := tValue.Field(i).Interface() // 获取属性的值
		fmt.Printf("%s %v = %v\n", field.Name, field.Type, value)
	}

	// 获取方法并执行方法
	println(tType.NumMethod())
	for i := 0; i < tType.NumMethod(); i++ {
		mType := tType.Method(i)
		fmt.Printf("%v %v\n", mType.Name, mType.Type)
		f := tValue.MethodByName(mType.Name)
		// 判断方法是否有效
		if f.IsValid() {
			// 调用无参方法
			if mType.Name == "UserCall" {
				f.Call(nil)
			} else {
				// 创建传入的参数
				params := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf("hello")}
				f.Call(params)
			}
		} else {
			fmt.Printf("method %v is not found\n", mType.Name)
		}
	}

}

func main() {
	var a float64 = 3.14
	aType := reflect.TypeOf(a)
	fmt.Println(aType)
	aValue := reflect.ValueOf(a)
	fmt.Println(aValue)
	u := User{
		Name: "jack",
		Age:  10,
	}
	doReflect(u)

}
