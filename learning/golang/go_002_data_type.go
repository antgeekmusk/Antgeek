package main

import (
	"fmt"
	"unsafe"
)

// 数据类型
func main() {
	//boolType()
	digitalType()
}

// 布尔型
func boolType() {
	var a bool
	// 默认是false
	fmt.Println(a)
	var b bool = true
	fmt.Println(b)
}

// 数字类型
func digitalType() {
	// ===================== 无符号整形 =====================

	// 无符号 8 位整型 (0 到 255)
	// 空间大小 :  1字节
	var a, b uint8 = 0, 255
	fmt.Println("Range of uint8:", a, b)
	fmt.Println("Size of uint8:", unsafe.Sizeof(a))

	// 无符号 16 位整型 (0 到 65535)
	// 空间大小 :  2字节
	var c, d uint16 = 0, 65535
	fmt.Println("Range of uint16:", c, d)
	fmt.Println("Size of uint16:", unsafe.Sizeof(c))

	// 无符号 32 位整型 (0 到 4294967295)
	// 空间大小 :  4字节
	var e, f uint32 = 0, 4294967295
	fmt.Println("Range of uint32:", e, f)
	fmt.Println("Size of uint32:", unsafe.Sizeof(e))

	// 无符号 64 位整型 (0 到 18446744073709551615)
	// 空间大小 :  8字节
	var g, h uint64 = 0, 18446744073709551615
	fmt.Println("Range of uint64:", g, h)
	fmt.Println("Size of uint64:", unsafe.Sizeof(g))

	// ===================== 有符号整形 =====================

	// 有符号 8 位整型 (-128 到 127)
	// 空间大小 :  1字节
	var a1, b1 int8 = -128, 127
	fmt.Println("Range of int8:", a1, b1)
	fmt.Println("Size of int8:", unsafe.Sizeof(a1))

	// 有符号 16 位整型 (-32768 到 32767)
	// 空间大小 :  2字节
	var c1, d1 int16 = -32768, 32767
	fmt.Println("Range of int16:", c1, d1)
	fmt.Println("Size of int16:", unsafe.Sizeof(c1))

	// 有符号 32 位整型 (-2147483648 到 2147483647)
	// 空间大小 :  4字节
	var e1, f1 int32 = -2147483648, 2147483647
	fmt.Println("Range of int32:", e1, f1)
	fmt.Println("Size of int32:", unsafe.Sizeof(e1))

	// 有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)
	// 空间大小 :  8字节
	var g1, h1 int64 = -9223372036854775808, -9223372036854775807
	fmt.Println("Range of int64:", g1, h1)
	fmt.Println("Size of int64:", unsafe.Sizeof(g1))

	// ===================== 浮点型 =====================

	// IEEE-754 32位浮点型数
	// 空间大小 :  4字节
	var a2 float32 = 3.14
	fmt.Println("Size of float32:", unsafe.Sizeof(a2))

	// IEEE-754 64位浮点型数
	// 空间大小 : 8字节
	var b2 float64 = 3.14
	fmt.Println("Size of float64:", unsafe.Sizeof(b2))

	// 32 位实数和虚数
	// 空间大小 : 8字节
	var c2 complex64 = 3.14
	var c3 complex64 = 2 + 2i
	var c4 complex64 = complex(1, 2)
	fmt.Println(c3)
	fmt.Println(c4)
	fmt.Println("复数的实部 : ", real(c4))
	fmt.Println("复数的虚部 : ", imag(c4))
	c5 := c4 + c3
	fmt.Println("复数加法", c5)
	c6 := c4 * c3
	fmt.Println("复数乘法", c6)
	fmt.Println("Size of complex64:", unsafe.Sizeof(c2))

	// 64 位实数和虚数
	// 空间大小 : 16字节
	var d2 complex128 = 3.14
	fmt.Println("Size of complex128:", unsafe.Sizeof(d2))

	// ===================== 其他数字类型 =====================

	// byte 类似 uint8
	// 空间大小 : 1字节
	var a3, b3 byte = 0, 255
	fmt.Println("Range of byte:", a3, b3)
	fmt.Println("Size of byte:", unsafe.Sizeof(a3))

	// rune 类似 类似 int32
	// 空间大小 : 4字节
	var c33, d3 rune = -2147483648, 2147483647
	fmt.Println("Range of rune:", c33, d3)
	fmt.Println("Size of rune:", unsafe.Sizeof(c33))

	// 32 或 64 位 (32位系统uint 是32,64位系统 uint 是64)
	// 空间大小 : 4字节(32位系统)/8字节(64位系统)
	var i uint = 1
	fmt.Println("Size of uint:", unsafe.Sizeof(i))

	// 32 或 64 位 (32位系统int 是32,64位系统 int 是64)
	// 空间大小 : 4字节(32位系统)/8字节(64位系统)
	var i1 int = 1
	fmt.Println("Size of int:", unsafe.Sizeof(i1))

	// 无符号整型，用于存放一个指针
	// 空间大小 : 4字节(32位系统)/8字节(64位系统)
	var x uintptr = 0
	fmt.Println("Size of uintptr:", unsafe.Sizeof(x))

}
