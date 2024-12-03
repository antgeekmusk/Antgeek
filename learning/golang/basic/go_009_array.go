package main

import "fmt"

// 数组
func main() {
	//arrayDemo1()
	//arrayDemo2()
	//arrayDemo3()
	arrayDemo4()

}

// 数组声明及访问
func arrayDemo1() {
	// 声明数组但是不初始化
	var nums [5]int
	// 遍历数组并进行初始化
	for i := 0; i < len(nums); i++ {
		nums[i] = i * 10
	}
	// 输出数组
	for _, n := range nums {
		println(n)
	}

	// 使用初始化列表来声明
	var nums2 = [3]int{1, 2, 3}
	println("使用初始化列表来声明", len(nums2))

	// 长度不确定可以用 ... 来确定,编译器会自动推断数组长度
	nums3 := [...]int{1, 2, 3, 5, 6, 4, 10}
	println("长度不确定可以用 ... 来确定,编译器会自动推断数组长度", len(nums3))

	// 通过下标来初始化元素
	nums4 := [8]int{1: 20, 5: 30}
	println("通过下标来初始化元素", len(nums4), nums4[1], nums4[5])
	nums5 := [...]int{1: 20, 5: 30}
	println("通过下标来初始化元素", len(nums5))
}

// 多维数组
func arrayDemo2() {
	// 声明
	nums := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	// 遍历
	for _, arr := range nums {
		for _, num := range arr {
			print(num, " ")
		}
		println()
	}

	// 创建一个空的二维数组
	arrs := [][]int{}

	row1 := []int{1, 2, 3}
	row2 := []int{7, 8, 9}
	arrs = append(arrs, row1)
	arrs = append(arrs, row2)

	// 打印
	fmt.Println(arrs[0])
	fmt.Println(arrs[1])

}

// 向函数传递数组
func arrayDemo3() {
	arr := [5]int{1, 2, 3, 4, 5}
	arr2 := []int{1, 2, 3, 4, 5}
	println(getAvg1(arr))
	println(getAvg2(arr2))
}

func getAvg1(arr [5]int) float32 {
	var sum float32 = 0
	for _, num := range arr {
		sum += float32(num)
	}
	return sum / float32(len(arr))
}

func getAvg2(arr []int) float32 {
	var sum float32 = 0
	for _, num := range arr {
		sum += float32(num)
	}
	return sum / float32(len(arr))
}

// 向函数传递数组指针
func arrayDemo4() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("调用函数之前", arr)
	arrElementAddOne(&arr)
	fmt.Println("调用函数之后", arr)
}
func arrElementAddOne(arr *[5]int) {
	for i := range *arr {
		(*arr)[i] = (*arr)[i] + 1
	}
}
