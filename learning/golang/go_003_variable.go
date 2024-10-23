package main

import "fmt"

var (
	l1 int
	l2 bool
	l3 string = "123"
)

// 变量
func main() {
	// 一次声明多个变量
	var a, b = 1, "hello"
	fmt.Println(a, b)

	// 数值类型(包括complex64/128)默认为 0
	var c int
	fmt.Println(c)
	// bool类型默认值为 false
	var d bool
	fmt.Println(d)
	// string类型默认为 ""
	var e string
	fmt.Println(e)
	// 以下类型为 nil
	var a1 *int
	var a2 []int
	var a3 map[string]int
	var a4 chan int
	var a5 func(string) int
	var a6 error // error 是接口
	if a2 == nil {
		fmt.Println("a2", a2)
	}
	if a3 == nil {
		fmt.Println("a3", a3)
	}
	fmt.Println(a1, a2, a3, a4, a5, a6)

	// 根据值类型自动推断
	var f = "qwe"
	fmt.Println(f)

	// 如果变量已经用var申明过,再用 := 会报错
	g := false // 相当于 var g bool = false
	fmt.Println(g)

	// 多变量声明

	// 类型相同的多个变量
	var h1, h2, h3 int
	h1, h2, h3 = 1, 2, 3
	fmt.Println(h1, h2, h3)

	// 类型不同的多个变量
	var i1, i2, i3 = 1, true, "123"
	fmt.Println(i1, i2, i3)

	// 用 := 声明
	j1, j2, j3 := 2, false, "aaa"
	fmt.Println(j1, j2, j3)

	// 这种因式分解类型的申明一般用于全局变量
	var (
		k1 int
		k2 string
	)
	k1 = 1
	k2 = "123"
	fmt.Println(k1, k2)
	fmt.Println(l1, l2, l3)

	// m1 赋值给 m2 ,是直接把值深拷贝了一份
	m1 := 1
	var m2 = m1
	fmt.Println("值地址", &m1, &m2)

	// string 也是基本类型,不是引用类型
	n1 := "hello"
	var n2 = n1
	fmt.Println("值地址", &n1, &n2)

	// 引用类型的浅拷贝
	o1 := []int{1, 2, 3, 4, 5} // 这是一个切片
	fmt.Println(o1)
	o2 := o1
	o1[0] = 100
	fmt.Printf("引用类型的浅拷贝地址一样%p %p", o1, o2)
	fmt.Println()
	fmt.Println(o1)
	fmt.Println(o2)

	// 交换变量
	p1, p2 := 1, 2
	fmt.Println("交换前", p1, p2)
	p1, p2 = p2, p1
	fmt.Println("交换后", p1, p2)

	// _ 的妙用,它可以将返回值中没用的值抛弃
	_, q1 := 3, 4
	_, q2, _, q3 := 1, 2, 3, 4
	fmt.Println(q1)
	fmt.Println(q2, q3)

}
