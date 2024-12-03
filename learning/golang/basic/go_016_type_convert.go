package main

import (
	"fmt"
	"strconv"
)

// 类型转换
func main() {
	tcDemo1()
	tcDemo2()
	tcDemo3()
	tcDemo4()
}

// 数值类型转换
func tcDemo1() {
	// 整型转浮点型
	var a int = 10
	var b float32 = float32(a)
	println(b)

	// 将浮点转整型
	var c float64 = 255
	var d uint8 = uint8(c)
	println(d)

	// 转换的时候要主要范围大小,超出的话会异常,但是不会报错
	var c1 float64 = -1
	var d1 uint8 = uint8(c1)
	println(d1)
}

// 字符串类型转换
func tcDemo2() {
	// 字符转整型
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num)
	}

	// 整数转为字符串
	str2 := strconv.Itoa(num)
	fmt.Println(str2)

	// 字符串转浮点
	s1 := "3.14"
	n1, _ := strconv.ParseFloat(s1, 64)
	fmt.Println(n1)

	// 浮点转字符串
	s2 := strconv.FormatFloat(n1, 'f', 2, 64)
	fmt.Println(s2)

}

// 类型断言
/*
语法 :
value.(type)
或者
value.(T)
*/
func tcDemo3() {
	var i interface{} = "hello"
	str, ok := i.(string)
	if ok {
		fmt.Println(str)
	} else {
		fmt.Println("转换失败")
	}
}

// 空接口类型 (空接口类型可以存储任何类型的值,实际应用中空接口经常被用来处理多种类型的值)
func printValue(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Println("int", t)
	case string:
		fmt.Println("string", t)
	default:
		fmt.Println("未知类型", t)
	}
}

func tcDemo4() {
	printValue(10)
	printValue("123")
	printValue(1.0)
}
