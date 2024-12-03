package main

import "fmt"

// 多个defer执行顺序
func multiDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

// return 和 defer 执行顺序
func returnFunc() (a, b int) {
	fmt.Println("return run")
	return a, b
}
func returnAndDefer() (int, int) {
	defer fmt.Println("defer run")
	return returnFunc()
}

func main() {
	multiDefer()
	// 执行结果
	// 3
	// 2
	// 1
	returnAndDefer()
	// 执行结果
	// return run
	// defer run
}
