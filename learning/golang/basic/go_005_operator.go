package main

import "fmt"

// 操作符
func main() {
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	fmt.Printf("a 变量的类型 %T\n", a)
	fmt.Printf("b 变量的类型 %T\n", b)
	fmt.Printf("c 变量的类型 %T\n", c)

	ptr = &a
	fmt.Printf("a的值为 %d\n", a)
	fmt.Printf("a的地址为 %p\n", &a)
	fmt.Printf("ptr的值就是a的地址 %p\n", ptr)
	fmt.Printf("*ptr 为 %d\n", *ptr)
}
