package main

import "fmt"

// 指针
func main() {
	//pointerDemo1()
	//pointerDemo2()
	//pointerDemo3()
	pointerDemo4()

}

// 简单使用
func pointerDemo1() {
	var a int
	var p *int
	a = 10
	p = &a
	fmt.Printf("a 的值 : %d\n", a)
	fmt.Printf("a 的地址 : %p\n", &a)
	fmt.Printf("p 的值 : %p\n", p)
	fmt.Printf("p 所指向的值 : %d\n", *p)
}

// 空指针
func pointerDemo2() {
	var a *int
	fmt.Printf("a的值 : %p\n", a)

	// 空指针判断
	if a == nil {
		println("a 是空指针")
	}
}

// 指针数组
func pointerDemo3() {
	// 定义一个普通的数组
	nums := [3]int{1, 2, 3}

	// 定义一个指针数组将nums中的元素都存储起来
	ps := [3]*int{}

	// 遍历数组,将地址存储到ps中
	for i, n := range nums {
		ps[i] = &n
	}

	// 遍历指针数组获取数组中的值
	for i, p := range ps {
		fmt.Printf("指针数组中的第 %d 个值为 : %d\n", i, *p)
	}
}

// 指向指针的指针(二级指针)
func pointerDemo4() {
	var a int = 10
	var p *int = &a
	var p2 **int = &p

	fmt.Printf("a : %d\n", a)
	fmt.Printf("p : %d\n", *p)
	fmt.Printf("p2 : %d\n", **p2)
}
