package main

import "fmt"

// 语法基础

func main() {
	//strConnect()
	//spaceStyle()
	testSprintf()
	testPrintf()
}

// 字符串连接
func strConnect() {
	fmt.Println("hello" + "Angteek!")
}

// 空格风格
func spaceStyle() {
	// 变量赋值用空格
	var a = 1
	fmt.Println(a)
	// 运算符之间用空格
	var b = a + 1
	fmt.Println(b)
	// 关键字和表达式之间用空格
	if b > 0 {
		fmt.Println(b)
	}
	// 在函数调用时，函数名和左边等号之间要使用空格，参数之间也要使用空格。
	c := add(a, b)
	fmt.Println(c)
}
func add(a, b int) int {
	return a + b
}

// 格式化字符串
func testSprintf() {
	// %d 是数字 %s是字符串
	var a = 666
	var b = "antgeek"
	var c = "a=%d & b=%s"
	var res = fmt.Sprintf(c, a, b)
	fmt.Println(res)
}

func testPrintf() {
	// %d 是数字 %s是字符串
	var a = 666
	var b = "antgeek"
	var c = "a=%d & b=%s"
	fmt.Printf(c, a, b)
}
