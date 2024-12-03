package main

import "fmt"

// 切片
func main() {
	sliceDemo1()
	sliceDemo2()
	sliceDemo3()
	sliceDemo4()
}

// 定义切片
func sliceDemo1() {
	// 数组不指定长度的方式
	var s1 []int
	printSliceInfo(s1)

	// 切片在未初始化之前是nil
	if s1 == nil {
		println("s1 是nil")
	}

	// 使用make 来创建切片
	var s2 []int = make([]int, 3)
	// 简写
	s3 := make([]int, 3)
	printSliceInfo(s2)
	printSliceInfo(s3)

	// 指定容量(容量是可选擦拭能)
	s4 := make([]int, 3, 6)
	printSliceInfo(s4)

}

// 切片初始化
func sliceDemo2() {
	// 直接初始化
	s1 := []int{1, 2, 3}
	printSliceInfo(s1)

	// 通过数组进行初始化
	arr1 := [6]int{1, 2, 3, 4, 5, 6}
	// start:end 如果start,end都不写就是整个数组
	s2 := arr1[:]
	printSliceInfo(s2)
	// 从数组下标下标2 到下标3 end的值是开区间
	s3 := arr1[2:4]
	printSliceInfo(s3)
	// 从数组下标下标2 到最后
	s4 := arr1[2:]
	printSliceInfo(s4)
	// 从数组下标下标3 到最前面
	s5 := arr1[:3]
	printSliceInfo(s5)

	// 通过切片初始化切片
	s6 := s5[:1]
	printSliceInfo(s6)

}

// append 和 copy 函数
func sliceDemo3() {
	var s1 []int
	printSliceInfo2(&s1)

	// 向切片添加元素 如果容量不够了之后就会自动扩容,1024之前扩容是2倍,之后是1/4
	s1 = append(s1, 0)
	printSliceInfo2(&s1)
	s1 = append(s1, 1)
	printSliceInfo2(&s1)
	s1 = append(s1, 2)
	printSliceInfo2(&s1)
	s1 = append(s1, 3)
	printSliceInfo2(&s1)
	s1 = append(s1, 4)
	printSliceInfo2(&s1)
	s1 = append(s1, 5)
	printSliceInfo2(&s1)

	// copy 函数
	s2 := make([]int, len(s1), cap(s1)*2)
	printSliceInfo2(&s2)
}

func printSliceInfo(x []int) {
	// 切片的长度可以通过len获取
	// 切片的容量可以通过cap获取
	fmt.Printf("切片的地址 %p,长度 %d,容量 %d,值 %v\n", &x, len(x), cap(x), x)
}

func printSliceInfo2(x *[]int) {
	// 切片的长度可以通过len获取
	// 切片的容量可以通过cap获取
	fmt.Printf("切片的地址 %p,长度 %d,容量 %d,值 %v\n", x, len(*x), cap(*x), *x)
}

// 切片扩容后地址会变吗
func sliceDemo4() {
	s1 := []int{1, 2, 3}
	fmt.Printf("容量 : %d,扩容前地址 : %p\n", cap(s1), &s1)
	s1 = append(s1, 4)
	fmt.Printf("容量 : %d,扩容后地址 : %p\n", cap(s1), &s1)
}
