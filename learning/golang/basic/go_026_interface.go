package main

import "fmt"

// 空接口可以用来标识任何类型

func myFunc(args interface{}) {
	fmt.Println("myFunc run")
	fmt.Println(args)

	// 判断传入的类型 (类型断言)
	value, ok := args.(string)
	if !ok {
		fmt.Println("args is not string")
	} else {
		fmt.Println("args is string,value = ", value)
	}
}

// 定义一个结构体
type A struct {
	name string
}

func main() {
	myFunc(A{name: "jack"})
	myFunc(10)
	myFunc("abc")
	myFunc(2.1)
}
