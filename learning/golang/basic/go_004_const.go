package main

import "unsafe"

// 常量赋值可以使用内置函数
const (
	i = "abcdefffffffffffff"
	j = len(i)
	k = unsafe.Sizeof(i)
)

func main() {
	// 常量不使用也可以
	const LENGTH int = 10
	const WIDTH int = 5

	// 赋值多个常量
	const a, b, c = 1, false, "abc"
	println(a, b, c)
	println(i, j, k)

	// iota 用法
	const (
		x1  = iota // 0
		x2         // 1
		x3         // 2
		x33 = "aa" // aa iota=3
		x4         // aa iota=4
		x5  = 100  // 100 iota=5
		x6         // 100 iota=6
		x7  = iota // 7
		x8         // 8
	)
	println(x1, x2, x3, x33, x4, x5, x6, x7, x8)
}
