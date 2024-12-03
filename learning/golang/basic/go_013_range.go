package main

import "fmt"

// range 用法
func main() {
	// 遍历数组
	arr := []int{1, 2, 3, 4}
	for index, num := range arr {
		println("遍历数组", index, num)
	}

	// 遍历切片
	s := []int{1, 2, 3, 4, 5}
	for index, num := range s {
		println("遍历切片", index, num)
	}

	// 遍历通道
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)

	for v := range ch {
		println("遍历通道", v)
	}

	// 遍历map
	m := map[string]string{"k1": "v1", "k2": "v2"}
	for k, v := range m {
		println("遍历map", k, v)
	}

	// 遍历字符串
	str := "hello"
	for i, c := range str {
		fmt.Printf("遍历字符串 %d %c\n", i, c)
	}

	// 其中 key 或者value 是可以缺省的

	// 缺省 value
	for index := range arr {
		println("缺省 value", index)
	}

	// 缺省 key
	for _, value := range arr {
		println("缺省 key", value)
	}

}
