package main

import "fmt"

// map
func main() {
	// 定义一个空的map
	m1 := make(map[string]int)
	fmt.Printf("m1 : 值 %v ; 长度 %d\n", m1, len(m1))

	// 定义一个容量为10 的 map
	m2 := make(map[string]int, 10)
	fmt.Printf("m2 : 值 %v ; 长度 %d\n", m2, len(m2))

	// 使用字面量创建map
	m3 := map[string]int{
		"tom":  10,
		"jack": 20,
	}
	fmt.Printf("m3 : 值 %v ; 长度 %d\n", m3, len(m3))

	// 根据key 获取value
	println(m3["tom"])

	// 如果没有 ok 是false
	v, ok := m3["aaa"]
	println(v, ok)

	// 修改元素
	m3["tom"] = 100
	fmt.Printf("m3 : 值 %v ; 长度 %d\n", m3, len(m3))

	// 删除原始
	delete(m3, "tom")
	fmt.Printf("删除元素 %v", m3)

}
